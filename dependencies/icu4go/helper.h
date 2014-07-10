#ifndef __HELPER_H__
#define __HELPER_H__

#include <unicode/utypes.h>
#include <unicode/ucsdet.h>

const char* detectCharset(void *detector, void *input, int input_len, int *status);


#endif //__HELPER_H__
