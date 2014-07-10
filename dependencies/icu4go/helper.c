#include "helper.h"
#include <string.h>
#include <unicode/utypes.h>
#include <unicode/ucsdet.h>

const char* detectCharset(void *detector, void *input, int input_len, int *status) {
	const UCharsetMatch *bestGuess;
	const char *bestGuessedCharset = NULL;
	ucsdet_setText((UCharsetDetector*)detector, (char*)input, input_len, status);
	if (*status != U_ZERO_ERROR) {
		return NULL;
	}
	bestGuess = ucsdet_detect((UCharsetDetector*)detector, status);
	if (*status != U_ZERO_ERROR) {
		return NULL;
	}
	if (bestGuess == 0) {
		return NULL;
	}
	bestGuessedCharset = ucsdet_getName(bestGuess, status);
	if (*status != U_ZERO_ERROR) {
		return NULL;
	}
	return bestGuessedCharset;
}
