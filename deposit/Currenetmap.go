package concurrent

import "image"

var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	return nil
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}
