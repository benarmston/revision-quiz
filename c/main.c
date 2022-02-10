/* main.c */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <libgen.h>
#include <errno.h>
#include <string.h>
#include <getopt.h>
#include <sys/types.h>
#include <stdint.h>

#define OPTSTR "vi:o:f:h"
#define USAGE_FMT "%s [-v] [-f hexflag] [-i input file] [-o output file] [-h] \n"
#define ERR_FOPEN_INPUT "fopen(input, r)"
#define ERR_FOPEN_OUTPUT "fopen(output, w)"
#define ERR_DO_THE_NEEDFUL "do_the_needful blew up"
#define DEFAULT_PROGNAME "george"

extern int errno;
extern char *optarg;
extern int opterr, optind;

typedef struct {
   int	    verbose;
   uint32_t flags;
   FILE	    *input;
   FILE	    *output;
} options_t;

int dumb_global_variable = -11;

void usage(char *progname, int opt);
int do_the_needful(options_t *options);
// int quiz();
// int version();

int main(int argc, char *argv[]) {


  printf("argc=%d\n", argc); /* debugging */
  int i;
  for (i = 0; i < argc; i++) {
    puts(argv[i]);
  }

   int opt;
   options_t options = {0, 0x0, stdin, stdout};

   opterr = 0;

   while ((opt = getopt(argc, argv, OPTSTR)) != EOF) {
      switch(opt) {
	 case 'i':
	    printf("got i\n");
	    if (!(options.input = fopen(optarg, "r")) ) {
	       perror(ERR_FOPEN_INPUT);
	       exit(EXIT_FAILURE);
	    }
	    break;

	 case 'o':
	    printf("got o\n");
	    if (!(options.output = fopen(optarg, "w")) ) {
	       perror(ERR_FOPEN_OUTPUT);
	       exit(EXIT_FAILURE);
	    }
	    break;

	 case 'f':
	    printf("got f\n");
	    options.flags = (uint32_t)strtoul(optarg, NULL, 16);
	    break;

	 case 'v':
	    printf("got v\n");
	    options.verbose += 1;
	    break;

	 case 'h':
	    printf("got h\n");
	 default:
	    printf("=> Unkown option %d = %c \n", opt, opt);
	    usage(basename(argv[0]), opt);
	    break;
      }
      if (do_the_needful(&options) != EXIT_SUCCESS) {
	 perror(ERR_DO_THE_NEEDFUL);
	 exit(EXIT_FAILURE);
      }
      return EXIT_SUCCESS;
   }
   return EXIT_SUCCESS;
}

void usage(char *progname, int opt) {
   if (opt != 'h') {
      printf("Unkown option %d = %c \n", opt, opt);
   }
   fprintf(stderr, USAGE_FMT, progname?progname:DEFAULT_PROGNAME);
   exit(EXIT_FAILURE);
   /* NOTREACHED */
}

int do_the_needful(options_t *options) {
   if (!options) {
      errno = EINVAL;
      return EXIT_FAILURE;
   }
   if (!options->input || !options->output) {
      errno = ENOENT;
      return EXIT_FAILURE;
   }
   /* Do the needful*/
   return EXIT_SUCCESS;
}

int version() {
   printf("Version 0.4.0 \n");
   printf("This project is still in active development \n");
   printf("Visit https://github.com/crus4d3/revision-quiz for the latest version \n");
   return 0;
}

int quiz() {
   char response [] = "";
   int score = 0;
   printf("Please enter all answers without any trialling spaces and as concisely as possible. \n");
   printf("Type 'quit' to quit at any time. \n");
   return 0;
}
