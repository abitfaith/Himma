package nsenter

/*
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *himma_pid;
	himma_pid = getenv("himma_pid");
	if (himma_pid) {
		//fprintf(stdout, "got himma_pid=%s\n", himma_pid);
	} else {
		//fprintf(stdout, "missing himma_pid env skip nsenter");
		return;
	}
	char *himma_cmd;
	himma_cmd = getenv("himma_cmd");
	if (himma_cmd) {
		//fprintf(stdout, "got himma_cmd=%s\n", himma_cmd);
	} else {
		//fprintf(stdout, "missing himma_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", himma_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(himma_cmd);
	exit(0);
	return;
}
*/
import "C"
