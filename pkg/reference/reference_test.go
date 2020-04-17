/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    htcp://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reference

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/resource/fake"
	"github.com/crossplane/crossplane-runtime/pkg/test"
)

// TODO(negz): Find a better home for this. It can't currently live alongside
// its contemporaries in pkg/resource/fake because it would cause an import
// cycle.
type FakeManagedList struct {
	runtime.Object

	Items []resource.Managed
}

func (fml *FakeManagedList) GetItems() []resource.Managed {
	return fml.Items
}

func TestToAndFromPtr(t *testing.T) {
	cases := map[string]struct {
		want string
	}{
		"Zero":    {want: ""},
		"NonZero": {want: "pointy"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := FromPtrValue(ToPtrValue(tc.want))
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("FromPtrValue(ToPtrValue(%s): -want, +got: %s", tc.want, diff)

			}
		})

	}
}

func TestResolve(t *testing.T) {
	errBoom := errors.New("boom")
	now := metav1.Now()
	value := "coolv"
	ref := &v1alpha1.Reference{Name: "cool"}

	controlled := &fake.Managed{}
	controlled.SetName(value)
	meta.SetExternalName(controlled, value)
	meta.AddControllerReference(controlled, meta.AsController(&corev1.ObjectReference{UID: types.UID("very-unique")}))

	type args struct {
		ctx context.Context
		req ResolutionRequest
	}
	type want struct {
		rsp ResolutionResponse
		err error
	}
	cases := map[string]struct {
		reason string
		c      client.Reader
		from   resource.Managed
		args   args
		want   want
	}{
		"FromDeleted": {
			reason: "Should return early if the referencing managed resource was deleted",
			from:   &fake.Managed{ObjectMeta: metav1.ObjectMeta{DeletionTimestamp: &now}},
			args: args{
				req: ResolutionRequest{},
			},
			want: want{
				rsp: ResolutionResponse{},
				err: nil,
			},
		},
		"AlreadyResolved": {
			reason: "Should return early if the current value is non-zero",
			from:   &fake.Managed{},
			args: args{
				req: ResolutionRequest{CurrentValue: value},
			},
			want: want{
				rsp: ResolutionResponse{ResolvedValue: value},
				err: nil,
			},
		},
		"Unresolvable": {
			reason: "Should return early if neither a reference or selector were provided",
			from:   &fake.Managed{},
			args: args{
				req: ResolutionRequest{},
			},
			want: want{
				err: nil,
			},
		},
		"GetError": {
			reason: "Should return errors encountered while getting the referenced resource",
			c: &test.MockClient{
				MockGet: test.NewMockGetFn(errBoom),
			},
			from: &fake.Managed{},
			args: args{
				req: ResolutionRequest{
					Reference: ref,
					To:        To{Managed: &fake.Managed{}},
					Extract:   ExternalName(),
				},
			},
			want: want{
				err: errors.Wrap(errBoom, errGetManaged),
			},
		},
		"SuccessfulResolve": {
			reason: "No error should be returned when the value is successfully extracted",
			c: &test.MockClient{
				MockGet: test.NewMockGetFn(nil, func(obj runtime.Object) error {
					meta.SetExternalName(obj.(metav1.Object), value)
					return nil
				}),
			},
			from: &fake.Managed{},
			args: args{
				req: ResolutionRequest{
					Reference: ref,
					To:        To{Managed: &fake.Managed{}},
					Extract:   ExternalName(),
				},
			},
			want: want{
				rsp: ResolutionResponse{
					ResolvedValue:     value,
					ResolvedReference: ref,
				},
			},
		},
		"ListError": {
			reason: "Should return errors encountered while listing potential referenced resources",
			c: &test.MockClient{
				MockList: test.NewMockListFn(errBoom),
			},
			from: &fake.Managed{},
			args: args{
				req: ResolutionRequest{
					Selector: &v1alpha1.Selector{},
				},
			},
			want: want{
				rsp: ResolutionResponse{},
				err: errors.Wrap(errBoom, errListManaged),
			},
		},
		"NoMatches": {
			reason: "Should return an error when no managed resources match the selector",
			c: &test.MockClient{
				MockList: test.NewMockListFn(nil),
			},
			from: &fake.Managed{},
			args: args{
				req: ResolutionRequest{
					Selector: &v1alpha1.Selector{},
					To:       To{List: &FakeManagedList{}},
				},
			},
			want: want{
				rsp: ResolutionResponse{},
				err: errors.New(errNoMatches),
			},
		},
		"SuccessfulSelect": {
			reason: "A managed resource with a matching controller reference should be selected and returned",
			c: &test.MockClient{
				MockList: test.NewMockListFn(nil),
			},
			from: controlled,
			args: args{
				req: ResolutionRequest{
					Selector: &v1alpha1.Selector{
						MatchControllerRef: func() *bool { t := true; return &t }(),
					},
					To: To{List: &FakeManagedList{Items: []resource.Managed{
						&fake.Managed{}, // A resource that does not match.
						controlled,      // A resource with a matching controller reference.
					}}},
					Extract: ExternalName(),
				},
			},
			want: want{
				rsp: ResolutionResponse{
					ResolvedValue:     value,
					ResolvedReference: &v1alpha1.Reference{Name: value},
				},
				err: nil,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := NewAPIResolver(tc.c, tc.from)
			got, err := r.Resolve(tc.args.ctx, tc.args.req)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nControllersMustMatch(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.rsp, got); diff != "" {
				t.Errorf("\n%s\nControllersMustMatch(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestControllersMustMatch(t *testing.T) {
	cases := map[string]struct {
		s    *v1alpha1.Selector
		want bool
	}{
		"NilSelector": {
			s:    nil,
			want: false,
		},
		"NilMatchControllerRef": {
			s:    &v1alpha1.Selector{},
			want: false,
		},
		"False": {
			s:    &v1alpha1.Selector{MatchControllerRef: func() *bool { f := false; return &f }()},
			want: false,
		},
		"True": {
			s:    &v1alpha1.Selector{MatchControllerRef: func() *bool { t := true; return &t }()},
			want: true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := ControllersMustMatch(tc.s)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ControllersMustMatch(...): -want, +got:\n%s", diff)
			}
		})
	}
}
