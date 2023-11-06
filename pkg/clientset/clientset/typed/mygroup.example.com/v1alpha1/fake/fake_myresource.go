// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/xuxant/myresource-crd/pkg/apis/mygroup.example.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyResources implements MyResourceInterface
type FakeMyResources struct {
	Fake *FakeMygroupV1alpha1
	ns   string
}

var myresourcesResource = v1alpha1.SchemeGroupVersion.WithResource("myresources")

var myresourcesKind = v1alpha1.SchemeGroupVersion.WithKind("MyResource")

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *FakeMyResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcesResource, c.ns, name), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *FakeMyResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MyResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcesResource, myresourcesKind, c.ns, opts), &v1alpha1.MyResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MyResourceList{ListMeta: obj.(*v1alpha1.MyResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.MyResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myResources.
func (c *FakeMyResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myresourcesResource, c.ns, opts))

}

// Create takes the representation of a myResource and creates it.  Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Create(ctx context.Context, myResource *v1alpha1.MyResource, opts v1.CreateOptions) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcesResource, c.ns, myResource), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Update(ctx context.Context, myResource *v1alpha1.MyResource, opts v1.UpdateOptions) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcesResource, c.ns, myResource), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *FakeMyResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(myresourcesResource, c.ns, name, opts), &v1alpha1.MyResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MyResourceList{})
	return err
}

// Patch applies the patch and returns the patched myResource.
func (c *FakeMyResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}
