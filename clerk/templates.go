package clerk

import "fmt"

type TemplatesService service

type Template struct {
	Object       string `json:"object"`
	Slug         string `json:"slug"`
	TemplateType string `json:"template_type"`
	Name         string `json:"name"`
	Position     int    `json:"position"`
	CanRevert    bool   `json:"can_revert"`
	CanDelete    bool   `json:"can_delete"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type TemplateExtended struct {
	*Template
	Subject            string   `json:"subject"`
	Markup             string   `json:"markup"`
	Body               string   `json:"body"`
	MandatoryVariables []string `json:"mandatory_variables"`
}

func (s *TemplatesService) ListAll(templateType string) ([]Template, error) {
	templateURL := fmt.Sprintf("%s/%s", TemplatesUrl, templateType)
	req, _ := s.client.NewRequest("GET", templateURL)

	var templates []Template

	_, err := s.client.Do(req, &templates)
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func (s *TemplatesService) Read(templateType, slug string) (*TemplateExtended, error) {
	templateURL := fmt.Sprintf("%s/%s/%s", TemplatesUrl, templateType, slug)

	req, _ := s.client.NewRequest("GET", templateURL)

	var templateExtended TemplateExtended

	_, err := s.client.Do(req, &templateExtended)
	if err != nil {
		return nil, err
	}

	return &templateExtended, nil
}

type UpsertTemplateRequest struct {
	Name               string   `json:"name"`
	Subject            string   `json:"subject,omitempty"`
	Markup             string   `json:"markup,omitempty"`
	Body               string   `json:"body"`
	MandatoryVariables []string `json:"mandatory_variables,omitempty"`
}

func (s *TemplatesService) Upsert(templateType, slug string, upsertTemplateRequest *UpsertTemplateRequest) (*TemplateExtended, error) {
	templateURL := fmt.Sprintf("%s/%s/%s", TemplatesUrl, templateType, slug)
	req, _ := s.client.NewRequest("PUT", templateURL, upsertTemplateRequest)

	var upsertedTemplate TemplateExtended

	_, err := s.client.Do(req, &upsertedTemplate)
	if err != nil {
		return nil, err
	}

	return &upsertedTemplate, nil
}

// Revert reverts a user template to the corresponding system template
func (s *TemplatesService) Revert(templateType, slug string) (*TemplateExtended, error) {
	templateURL := fmt.Sprintf("%s/%s/%s/revert", TemplatesUrl, templateType, slug)
	req, _ := s.client.NewRequest("POST", templateURL)

	var templateExtended TemplateExtended

	_, err := s.client.Do(req, &templateExtended)
	if err != nil {
		return nil, err
	}

	return &templateExtended, nil
}

// Delete deletes a custom user template
func (s *TemplatesService) Delete(templateType, slug string) (*DeleteResponse, error) {
	templateURL := fmt.Sprintf("%s/%s/%s", TemplatesUrl, templateType, slug)
	req, _ := s.client.NewRequest("DELETE", templateURL)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}