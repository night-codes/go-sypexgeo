/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "sp.go"

#include <stdlib.h>

#ifndef GO_SYPEGEO
#define GO_SYPEGEO

typedef struct {
	int id;
	float lat;
	float lon;
	char* name_en;
	char* name_ru;
} spxCity;

typedef struct {
	int id;
	char* iso;
	char* name_en;
	char* name_ru;
} spxRegion;

typedef struct {
	int id;
	char* iso;
	float lat;
	float lon;
	char* name_en;
	char* name_ru;
} spxCountry;

typedef struct {
	spxCity city;
	spxCountry country;
	spxRegion region;
	int err;
} spxInfo;

static void spxFree(spxInfo s) {
	free(s.city.name_en);
	free(s.city.name_ru);
	free(s.country.iso);
	free(s.country.name_en);
	free(s.country.name_ru);
	free(s.region.iso);
	free(s.region.name_en);
	free(s.region.name_ru);
}
#endif

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif

extern void spxConnect(char* c);
extern spxInfo spxGetCity(char* ip, int full);
extern char* spxGetCountry(char* ip);
extern int spxGetCountryID(char* ip);

#ifdef __cplusplus
}
#endif
