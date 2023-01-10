//go:build ignore

package main

import (
	"log"

	// "ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
			spec.AddPathItem("/events/{id}/state", ogen.NewPathItem().
				SetPatch(ogen.NewOperation().
					SetOperationID("patchState").
					AddTags("Event").
					AddResponse("200", ogen.NewResponse()),
				).
				AddParameters(
					ogen.NewParameter().
						InPath().
						SetName("id").
						SetRequired(true).
						SetSchema(ogen.Int()),
				),
			)
			return nil
		}),
		entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
			spec.AddPathItem("/events/{id}/participants", ogen.NewPathItem().
				SetPost(ogen.NewOperation().
					SetOperationID("postEventParticipants").
					AddTags("Event").
					AddResponse("200", ogen.NewResponse()),
				).
				AddParameters(
					ogen.NewParameter().
						InPath().
						SetName("id").
						SetRequired(true).
						SetSchema(ogen.Int()),
				),
			)
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	// ogent, err := ogent.NewExtension(spec)
	// if err != nil {
	// 	log.Fatalf("creating ogent extension: %v", err)
	// }
	// err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	// if err != nil {
	// 	log.Fatalf("running ent codegen: %v", err)
	// }
}
