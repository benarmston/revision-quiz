#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define OPTSTR "bEnsTv-:"

int longopts(char *optarg);

int main(int argc, char **argv) {
   int i;
   int option;

   /* parse short options */

   while ((option = getopt(argc, argv, OPTSTR)) != -1) {
      switch (option) {
      case 'b':
         puts("Put line numbers next to non-blank lines");
         break;
      case 'E':
         puts("Show the ends of lines as $");
         break;
      case 'n':
         puts("Put line numbers next to all lines");
         break;
      case 's':
         puts("Suppress printing repeated blank lines");
         break;
      case 'T':
         puts("Show tabs as ^I");
         break;
      case 'v':
         puts("Verbose");
         break;
      case '-':
	 printf("%s \n", optarg);
	 longopts(optarg);
	 break;
      default:                                       /* '?' */
         puts("usage: simple -bEnsTv");
         return EXIT_FAILURE;
         //puts("What's that??");
      }
   }

   /* print the rest of the command line */

   // puts("------------------------------");

   // for (i = 0; i < optind; i++) {
   //    puts(argv[i]);
   // }

   // puts("------------------------------");

   // for (i = optind; i < argc; i++) {
   //    puts(argv[i]);
   // }

   return 0;
}

int longopts(char *optarg) {
   printf("%s", optarg);
   if (*optarg == "version") {
      puts("Print version information \n");
      return EXIT_SUCCESS;
   }
   else {
      printf("Unkown option %s \n", optarg);
      return EXIT_FAILURE;
   }
}
