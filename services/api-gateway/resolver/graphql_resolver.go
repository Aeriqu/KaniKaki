// Package resolver contains all of the resolvers used to handle GraphQL
// requests.
package resolver

type RootResolver struct{}
type QueryResolver struct{}
type MutationResolver struct{}

func(root *RootResolver) Query() *QueryResolver {
  return &QueryResolver{}
}

func(root *RootResolver) Mutation() *MutationResolver {
  return &MutationResolver{}
}