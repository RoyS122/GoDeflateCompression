package main

import (
	"bytes"
	"fmt"
)

func main() {

	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales mauris arcu, at commodo eros accumsan ut. Donec ut sapien aliquet, lobortis purus in, lacinia metus. Duis sit amet pulvinar leo. Morbi blandit magna dolor, eu pretium dolor feugiat a. Integer non leo ipsum. Sed at mi gravida, lobortis felis eu, facilisis libero. Aliquam ac tellus sit amet ex commodo consequat. Nulla auctor venenatis ornare. Phasellus ultricies leo vel nisl congue, in lobortis libero posuere. Phasellus at est sit amet nibh euismod blandit. Phasellus quis pharetra dui. Vivamus fringilla erat et velit mollis, at vulputate risus iaculis. Donec ultrices quam eget mauris tristique consequat. Sed iaculis id dolor et feugiat. Duis sit amet massa interdum, lacinia ante quis, iaculis elit.

Nullam tempor lorem eget ante dictum, sed venenatis arcu tristique. Maecenas justo enim, egestas vitae orci ac, consectetur scelerisque ex. Phasellus vitae sem libero. Donec rutrum mollis tortor a vulputate. Praesent ullamcorper ultricies enim eu blandit. Cras efficitur turpis gravida venenatis elementum. Vivamus bibendum semper rutrum. Aenean aliquam, nunc in iaculis ultrices, eros lectus blandit turpis, ut porttitor ante leo ac sapien. Donec et urna enim. Nulla posuere odio eget dignissim luctus. Donec est massa, faucibus vitae orci ut, dictum pharetra felis. Etiam fringilla libero eget sem mattis hendrerit. Maecenas leo nibh, bibendum vitae consequat et, facilisis sed purus.

Donec at tortor magna. Etiam eget sem eget neque cursus viverra. Maecenas eu accumsan risus, a efficitur dolor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec maximus ornare mauris, vel tincidunt orci ultrices eget. Vestibulum ornare vitae massa vitae aliquam. Proin vitae orci quam. Vivamus purus diam, aliquam vel est eget, sodales rutrum turpis. Curabitur eu consectetur leo. Morbi sagittis dapibus nisi, id consequat arcu egestas at.

Cras efficitur in arcu sed fermentum. Suspendisse ex nibh, dictum nec elit et, posuere accumsan nibh. Pellentesque orci eros, vehicula ac neque ac, fringilla mattis risus. Nunc sed maximus est, et varius turpis. In hac habitasse platea dictumst. Curabitur lectus arcu, iaculis ut lacinia ut, ullamcorper sit amet odio. Proin lobortis, elit id posuere dapibus, dolor dui iaculis elit, vitae condimentum nunc turpis eu dui. Nulla viverra et eros sit amet tempus. Aenean quam neque, iaculis eu tincidunt nec, ullamcorper a nisi. Donec fermentum dui sed porttitor molestie. In sodales suscipit eros, eget sollicitudin leo aliquam in. Nulla non lectus ullamcorper, fermentum urna rutrum, gravida nulla. Donec viverra, leo eget tempus luctus, ligula augue ornare orci, eu sagittis diam turpis quis quam.

Phasellus feugiat, nibh at iaculis egestas, ex dolor fermentum diam, eget euismod lacus nunc ultrices ex. Proin maximus elit ac pulvinar congue. Donec scelerisque odio sed metus bibendum, lobortis tincidunt nunc mattis. Suspendisse interdum, diam id aliquet ultricies, ante mi pellentesque nisi, sit amet euismod nulla ipsum id nibh. Nam enim erat, congue a vulputate sollicitudin, tincidunt eget dolor. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Fusce posuere, sem sed venenatis tincidunt, lorem nunc egestas sapien, vitae placerat ligula massa vel massa. Praesent erat dui, interdum sed neque et, vehicula porttitor urna. Sed tincidunt neque sem, in porta odio efficitur euismod. Interdum et malesuada fames ac ante ipsum primis in faucibus.`

	bin, tree, totalChars, usedLZ := FullCompression(text)
	var treeAsSerialized bytes.Buffer

	// Décompression
	final := FullDecompression(bin, tree, totalChars, usedLZ)

	fmt.Println("Taille originale: ", len([]byte(text)))
	fmt.Println("Taille compressé", len(bin)+len(treeAsSerialized.Bytes())+1+8)

	if final == text {
		fmt.Println("\nSuccès : le texte décompressé est identique à l'original !")
	} else {
		fmt.Println("\nErreur : le texte décompressé est différent de l'original.")
	}

}
