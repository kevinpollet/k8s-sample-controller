/*
The MIT License (MIT)

Copyright (c) 2022 kevinpollet <pollet.kevin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kevinpollet/k8s-sample-controller/pkg/apis/sample/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJobs implements JobInterface
type FakeJobs struct {
	Fake *FakeSampleV1alpha1
	ns   string
}

var jobsResource = schema.GroupVersionResource{Group: "sample.io", Version: "v1alpha1", Resource: "jobs"}

var jobsKind = schema.GroupVersionKind{Group: "sample.io", Version: "v1alpha1", Kind: "Job"}

// Get takes name of the job, and returns the corresponding job object, and an error if there is any.
func (c *FakeJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jobsResource, c.ns, name), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// List takes label and field selectors, and returns the list of Jobs that match those selectors.
func (c *FakeJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jobsResource, jobsKind, c.ns, opts), &v1alpha1.JobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JobList{ListMeta: obj.(*v1alpha1.JobList).ListMeta}
	for _, item := range obj.(*v1alpha1.JobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *FakeJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jobsResource, c.ns, opts))

}

// Create takes the representation of a job and creates it.  Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Create(ctx context.Context, job *v1alpha1.Job, opts v1.CreateOptions) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jobsResource, c.ns, job), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// Update takes the representation of a job and updates it. Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Update(ctx context.Context, job *v1alpha1.Job, opts v1.UpdateOptions) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jobsResource, c.ns, job), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// Delete takes name of the job and deletes it. Returns an error if one occurs.
func (c *FakeJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(jobsResource, c.ns, name, opts), &v1alpha1.Job{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JobList{})
	return err
}

// Patch applies the patch and returns the patched job.
func (c *FakeJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}
