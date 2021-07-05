package payload

type IOSPayload map[string]interface{}

func (p *IOSPayload) Initial() {
	(*p)["aps"] = &APNs{
		ContentAvailable: 0,
	}
}

func (p *IOSPayload) GetAPNs() *APNs {
	return (*p)["aps"].(*APNs)
}

func (p *IOSPayload) AddExtra(key string, val string) *IOSPayload {
	if key == "aps" { // 防止自定义aps覆盖
		return p
	}
	(*p)[key] = val
	return p
}

type Alert struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Body     string `json:"body"`
}

type APNs struct {
	Alert            *Alert `json:"alert"`
	Badge            string `json:"badge,omitempty"`
	Sound            string `json:"sound"`
	ContentAvailable int    `json:"content-available,omitempty"`
	Category         string `json:"category,omitempty"`
	QFAttach         string `json:"QFAttach,omitempty"`
}

func (a *APNs) SetTitle(title string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.Title = title
	return a
}

func (a *APNs) SetSubTitle(subTitle string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.SubTitle = subTitle
	return a
}

func (a *APNs) SetBody(body string) *APNs {
	if a.Alert == nil {
		a.Alert = &Alert{}
	}
	a.Alert.Body = body
	return a
}

func (a *APNs) SetBadge(badge string) *APNs {
	a.Badge = badge
	return a
}

func (a *APNs) SetSound(sound string) *APNs {
	a.Sound = sound
	return a
}

func (a *APNs) SetContentAvailable(val int) *APNs {
	a.ContentAvailable = val
	return a
}

func (a *APNs) SetCategory(category string) *APNs {
	a.Category = category
	return a
}

func (a *APNs) SetImg(img string) *APNs {
	a.QFAttach = img
	return a
}
