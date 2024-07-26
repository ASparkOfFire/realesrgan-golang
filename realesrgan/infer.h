// infer.h
#include <stdlib.h>

#ifndef INFER_H
#define INFER_H

#ifdef __cplusplus
extern "C" {
#endif

// Define your parameters struct if needed
struct RealESRGanParams
{
    int gpuid;
    int tta;  // Use `int` instead of `bool` for C compatibility
    int scale;
    int tilesize;
    int prepadding;
    const char* model_path;
	const char* param_path;
};

struct Buffer {
    const unsigned char* data;
    size_t size;
};

// Function declarations

void* create_realesrgan_instance(int gpuid,
                                 int tta,
                                 int scale,
                                 int tilesize,
                                 int prepadding,
                                 const char* model_path,
                                 const char* param_path);
struct Buffer process_image(const unsigned char* input_image, size_t length, const char* image_format, void* predictor, int scale);
void free_buffer(struct Buffer buffer);
void delete_realesrgan(void *realesrgan);


#ifdef __cplusplus
}
#endif

#endif // INFER_H
