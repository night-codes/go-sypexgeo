#include <stdio.h>
#include "sp.h"


void printInfo(spxInfo info) {
	printf("Err %i\n", info.err);
	printf("City %i %f %f %s %s\n", info.city.id, info.city.lat, info.city.lon, info.city.name_en, info.city.name_ru);
	printf("Coutry %i %s %f %f %s %s\n", info.country.id, info.country.iso, info.country.lat, info.country.lon, info.country.name_en, info.country.name_ru);
	printf("Region %i %s %s %s\n-----\n", info.region.id, info.region.iso, info.region.name_en, info.region.name_ru);
}


int main() {
	spxConnect("../SxGeoCity.dat");
	printInfo(spxGetCity("93.73.35.74", 1));
	printInfo(spxGetCity("13.70.15.11", 0));

	printf("Coutry name: %s\n", spxGetCity("53.12.92.225", 0).country.name_en);
	printf("Coutry code: %s\n", spxGetCountry("53.12.92.225"));
	printf("Coutry ID: %i\n", spxGetCountryID("53.12.92.225"));
}
