package graph

import (
	"bipper-backend/models"
	"errors"
	"log"
	"reflect"
)

func ApplyCursorsToEdges(allEdges interface{}, before *string, after *string) ([]models.Node, error) {
	if reflect.TypeOf(allEdges).Kind() != reflect.Slice {
		return nil, errors.New("no slice model.Node")
	}
	s := reflect.ValueOf(allEdges)
	slice := make([]models.Node, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		f := s.Index(i).Interface().(models.Node)
		slice = append(slice, f)
	}

	//Initialize edges to be allEdges.
	edges := slice
	// 2 If after is set:
	if after != nil {
		// a Let afterEdge be the edge in edges whose cursor is equal to the after argument.
		afterEdge := ""
		afterEdgeIndex := 0

		for i := range edges {
			s := reflect.ValueOf(edges[i])
			if s.Kind() == reflect.Ptr {
				s = reflect.Indirect(s)
			}
			if s.Kind() != reflect.Struct {
				log.Fatal("unexpected type")
			}
			typeOfT := s.Type()
			var edgeID string
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if typeOfT.Field(i).Name == "ID" {
					edgeID = f.String()
				}
			}
			if edgeID == *after {
				afterEdge = *after
				afterEdgeIndex = i
				break
			}
		}

		// b If afterEdge exists:  Remove all elements of edges before and including afterEdge.
		if afterEdge != "" && afterEdgeIndex < len(edges) {
			edges = append(edges[afterEdgeIndex+1:])
		}
	}
	// 3 If before is set:
	if before != nil {
		// a Let beforeEdge be the edge in edges whose cursor is equal to the before argument.
		beforeEdge := ""
		beforeEdgeIndex := 0
		for i := range edges {
			s := reflect.ValueOf(edges[i])
			if s.Kind() == reflect.Ptr {
				s = reflect.Indirect(s)
			}
			if s.Kind() != reflect.Struct {
				log.Fatal("unexpected type")
			}
			typeOfT := s.Type()
			var edgeID string
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if typeOfT.Field(i).Name == "ID" {
					edgeID = f.String()
				}
			}
			if edgeID == *before {
				beforeEdge = *before
				beforeEdgeIndex = i
				break
			}
		}
		// b If beforeEdge exists: Remove all elements of edges after and including beforeEdge.
		if beforeEdge != "" && beforeEdgeIndex < len(edges) {
			edges = append(edges[:beforeEdgeIndex])
		}
	}
	//Return edges.
	return edges, nil
}

func EdgesToReturn(allEdges interface{}, before *string, after *string, first *int, last *int) ([]models.Node, error) {
	edges, err := ApplyCursorsToEdges(allEdges, before, after)
	//fmt.Printf("%v\n", edges)
	if err != nil {
		return nil, err
	}

	if first != nil {
		//If first is less than 0:
		//Throw an error.
		if *first < 0 {
			return nil, errors.New("first less than 0")
		}
		//	If edges has length greater than than first:
		//Slice edges to be of length first by removing edges from the end of edges.
		if len(edges) > *first {
			edges = append(edges[0:*first])
		}
	}
	if last != nil {
		//If last is less than 0:
		//Throw an error.
		if *last < 0 {
			return nil, errors.New("last less than 0")
		}
		//	If edges has length greater than than last:
		//Slice edges to be of length last by removing edges from the start of edges.
		if len(edges) > *last {
			edges = append(edges[len(edges)-*last:])
		}
	}

	return edges, nil
}

func HasPreviousPage(allEdges interface{}, before *string, after *string, first *int, last *int) (bool, error) {
	if reflect.TypeOf(allEdges).Kind() != reflect.Slice {
		return false, errors.New("no slice model.Node")
	}
	s := reflect.ValueOf(allEdges)
	slice := make([]models.Node, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		f := s.Index(i).Interface().(models.Node)
		slice = append(slice, f)
	}

	// 1 If last is set:
	if last != nil {
		// 1-a Let edges be the result of calling ApplyCursorsToEdges(allEdges, before, after).
		edges, err := ApplyCursorsToEdges(allEdges, before, after)
		if err != nil {
			return false, err
		}
		// 1-b If edges contains more than last elements return true, otherwise false.
		if len(edges) > *last {
			return true, nil
		}
	}
	// 2 If after is set:
	if after != nil {
		// 2-a If the server can efficiently determine that elements exist prior to after, return true.
		afterIndex := 0

		for i := range slice {
			s := reflect.ValueOf(slice[i])
			if s.Kind() == reflect.Ptr {
				s = reflect.Indirect(s)
			}
			if s.Kind() != reflect.Struct {
				log.Fatal("unexpected type")
			}
			typeOfT := s.Type()
			var allEdgeID string
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if typeOfT.Field(i).Name == "ID" {
					allEdgeID = f.String()
				}
			}
			if allEdgeID == *after {
				afterIndex = i
				break
			}
		}
		if afterIndex > 0 {
			return true, nil
		}
	}
	// 3 Return false.
	return false, nil
}

func HasNextPage(allEdges interface{}, before *string, after *string, first *int, last *int) (bool, error) {
	if reflect.TypeOf(allEdges).Kind() != reflect.Slice {
		return false, errors.New("no slice model.Node")
	}
	s := reflect.ValueOf(allEdges)
	slice := make([]models.Node, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		f := s.Index(i).Interface().(models.Node)
		slice = append(slice, f)
	}

	// 1 If first is set:
	if first != nil {
		// 1-a Let edges be the result of calling ApplyCursorsToEdges(allEdges, before, after).
		edges, err := ApplyCursorsToEdges(allEdges, before, after)
		if err != nil {
			return false, err
		}
		// 1-b If edges contains more than first elements return true, otherwise false.
		if len(edges) > *first {
			return true, nil
		}
	}
	// 2 If before is set:
	if before != nil {
		// 2-a If the server can efficiently determine that elements exist following before, return true
		beforeIndex := len(slice)
		for i := range slice {
			s := reflect.ValueOf(slice[i])
			if s.Kind() == reflect.Ptr {
				s = reflect.Indirect(s)
			}
			if s.Kind() != reflect.Struct {
				log.Fatal("unexpected type")
			}
			typeOfT := s.Type()
			var allEdgeID string
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if typeOfT.Field(i).Name == "ID" {
					allEdgeID = f.String()
				}
			}
			if allEdgeID == *before {
				beforeIndex = i
				break
			}
		}
		if beforeIndex < len(slice) {
			return true, nil
		}
	}
	// 3 Return false.
	return false, nil
}
