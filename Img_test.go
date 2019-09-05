package common

import (
	"fmt"
	"testing"
)

func TestImgBase64ToImg(t *testing.T) {
	b64 := `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQQAAAIzCAYAAADmn2eCAAAGXElEQVR4nO3ZwWrCQABF0YwGwYUQ8v+/lC8StLXQu1ZoiS3nrGb5NnNhmHH/NP0R27ZNy7JM67pOl8tlOh6P0xhj71nwbxz2HgC8D0EAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAMm/btveGlxwOh2mMsfcM+JfmZVn23vC0RwzO5/M0z7MowC+Y13Xde8PTHhF4xOB0On2dRQF+1rher/e9R7ziOwSCAD9v3G63PxWEByGA3zG7XMA3345ABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAIghABAGIIAARBCCCAEQQgAgCEEEAMm/btvcG4E2M+6e9RwDvwZMBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQAQBiCAAEQQgggBEEIAIAhBBACIIQD4ALKMc9P/c5S0AAAAASUVORK5CYII=`
	e := ImgBase64ToImg(b64, "static/1.png")
	t.Log(e)
}

func TestImgToBase64(t *testing.T) {
	e, s := ImgToBase64("static/1.png")
	fmt.Println(e, s)
}