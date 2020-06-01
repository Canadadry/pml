package renderer

import (
	"errors"
	"fmt"
	"pml/pkg/token"
)

const (
	itemDocument  = "Document"
	itemPage      = "Page"
	itemRectangle = "Rectangle"
	itemText      = "Text"
	itemFont      = "Font"
	itemImage     = "Image"
	itemVector    = "Vector"
)

var (
	errItemNotFound               = errors.New("errItemNotFound")
	errChildrenNotAllowed         = errors.New("errChildrenNotAllowed")
	errPropertyDefinitionNotFound = errors.New("errPropertyDefinitionNotFound")
)

type itemDefinition struct {
	authorizedChildren   []string
	authorizedProperties map[string]token.TokenType
}

type itemDefinitions struct {
	items map[string]itemDefinition
}

func (id itemDefinitions) validateChildType(itemName string, childType string) error {
	item, ok := id.items[itemName]
	if !ok {
		return fmt.Errorf("%w searching %s", errItemNotFound, itemName)
	}

	for _, authorizedChild := range item.authorizedChildren {
		if childType == authorizedChild {
			return nil
		}
	}
	return fmt.Errorf("in %s : %w children %s", itemName, errChildrenNotAllowed, childType)
}

func (id itemDefinitions) getPropertyType(itemName string, propertyName string) (token.TokenType, error) {
	item, ok := id.items[itemName]
	if !ok {
		return token.ILLEGAL, fmt.Errorf("%w searching %s", errItemNotFound, itemName)
	}

	pptType, ok := item.authorizedProperties[propertyName]
	if !ok {
		return token.ILLEGAL, fmt.Errorf("in %s : %w property %s", itemName, errPropertyDefinitionNotFound, propertyName)
	}

	return pptType, nil
}

var items = itemDefinitions{
	items: map[string]itemDefinition{
		itemDocument: itemDefinition{
			authorizedChildren:   []string{itemPage, itemFont},
			authorizedProperties: map[string]token.TokenType{},
		},
		itemFont: itemDefinition{
			authorizedChildren: []string{},
			authorizedProperties: map[string]token.TokenType{
				"file": token.STRING,
				"name": token.STRING,
			},
		},
		itemPage: itemDefinition{
			authorizedChildren:   []string{itemRectangle, itemText, itemImage, itemVector},
			authorizedProperties: map[string]token.TokenType{},
		},
		itemRectangle: itemDefinition{
			authorizedChildren: []string{itemRectangle, itemText, itemImage, itemVector},
			authorizedProperties: map[string]token.TokenType{
				"x":      token.FLOAT,
				"y":      token.FLOAT,
				"width":  token.FLOAT,
				"height": token.FLOAT,
				"color":  token.COLOR,
			},
		},
		itemText: itemDefinition{
			authorizedChildren: []string{},
			authorizedProperties: map[string]token.TokenType{
				"text":     token.STRING,
				"x":        token.FLOAT,
				"y":        token.FLOAT,
				"width":    token.FLOAT,
				"height":   token.FLOAT,
				"color":    token.COLOR,
				"align":    token.IDENTIFIER,
				"fontName": token.STRING,
				"fontSize": token.FLOAT,
			},
		},
		itemImage: itemDefinition{
			authorizedChildren: []string{},
			authorizedProperties: map[string]token.TokenType{
				"x":      token.FLOAT,
				"y":      token.FLOAT,
				"width":  token.FLOAT,
				"height": token.FLOAT,
				"file":   token.STRING,
			},
		},
		itemVector: itemDefinition{
			authorizedChildren: []string{},
			authorizedProperties: map[string]token.TokenType{
				"x":      token.FLOAT,
				"y":      token.FLOAT,
				"width":  token.FLOAT,
				"height": token.FLOAT,
				"file":   token.STRING,
			},
		},
	},
}
