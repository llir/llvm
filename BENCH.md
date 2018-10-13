# Parsing LLVM IR of coreutils

Testdata located at https://github.com/decomp/testdata/tree/master/coreutils/testdata

## No semantic actions

```
path: testdata/[.ll
took: 32.266058ms

path: testdata/b2sum.ll
took: 31.941161ms

path: testdata/base32.ll
took: 17.655541ms

path: testdata/base64.ll
took: 16.5282ms

path: testdata/basename.ll
took: 12.393106ms

path: testdata/cat.ll
took: 14.88439ms

path: testdata/chcon.ll
took: 35.236751ms

path: testdata/chgrp.ll
took: 33.20814ms

path: testdata/chmod.ll
took: 33.203252ms

path: testdata/chown.ll
took: 34.931635ms

path: testdata/chroot.ll
took: 17.991691ms

path: testdata/cksum.ll
took: 13.146118ms

path: testdata/comm.ll
took: 16.037409ms

path: testdata/cp.ll
took: 53.62337ms

path: testdata/csplit.ll
took: 22.05414ms

path: testdata/cut.ll
took: 19.213034ms

path: testdata/date.ll
took: 49.920084ms

path: testdata/dd.ll
took: 28.254343ms

path: testdata/df.ll
took: 49.528454ms

path: testdata/dir.ll
took: 89.729608ms

path: testdata/dircolors.ll
took: 20.158047ms

path: testdata/dirname.ll
took: 14.852603ms

path: testdata/du.ll
took: 80.229389ms

path: testdata/echo.ll
took: 12.867711ms

path: testdata/env.ll
took: 12.572114ms

path: testdata/expand.ll
took: 18.923754ms

path: testdata/expr.ll
took: 20.291402ms

path: testdata/factor.ll
took: 50.324439ms

path: testdata/false.ll
took: 12.3716ms

path: testdata/fmt.ll
took: 20.556029ms

path: testdata/fold.ll
took: 20.892052ms

path: testdata/getlimits.ll
took: 20.450926ms

path: testdata/ginstall.ll
took: 71.107709ms

path: testdata/groups.ll
took: 14.322809ms

path: testdata/head.ll
took: 21.334622ms

path: testdata/hostid.ll
took: 15.418029ms

path: testdata/id.ll
took: 19.33295ms

path: testdata/join.ll
took: 26.247752ms

path: testdata/kill.ll
took: 18.110894ms

path: testdata/link.ll
took: 15.169345ms

path: testdata/ln.ll
took: 39.532417ms

path: testdata/logname.ll
took: 12.275252ms

path: testdata/ls.ll
took: 94.382067ms

path: testdata/make-prime-list.ll
took: 1.178483ms

path: testdata/md5sum.ll
took: 19.858581ms

path: testdata/mkdir.ll
took: 21.705242ms

path: testdata/mkfifo.ll
took: 13.317481ms

path: testdata/mknod.ll
took: 20.180666ms

path: testdata/mktemp.ll
took: 18.656773ms

path: testdata/mv.ll
took: 79.127165ms

path: testdata/nice.ll
took: 14.025481ms

path: testdata/nl.ll
took: 16.722598ms

path: testdata/nohup.ll
took: 17.927111ms

path: testdata/nproc.ll
took: 16.193828ms

path: testdata/numfmt.ll
took: 31.02417ms

path: testdata/od.ll
took: 36.931837ms

path: testdata/paste.ll
took: 14.541513ms

path: testdata/pathchk.ll
took: 17.750618ms

path: testdata/pinky.ll
took: 21.074214ms

path: testdata/pr.ll
took: 52.460187ms

path: testdata/printenv.ll
took: 12.29439ms

path: testdata/printf.ll
took: 17.753412ms

path: testdata/ptx.ll
took: 39.672154ms

path: testdata/pwd.ll
took: 16.815041ms

path: testdata/readlink.ll
took: 26.14529ms

path: testdata/realpath.ll
took: 29.018492ms

path: testdata/rm.ll
took: 40.533879ms

path: testdata/rmdir.ll
took: 12.817569ms

path: testdata/runcon.ll
took: 14.087019ms

path: testdata/seq.ll
took: 21.206139ms

path: testdata/sha1sum.ll
took: 21.799649ms

path: testdata/sha224sum.ll
took: 26.261186ms

path: testdata/sha256sum.ll
took: 29.444181ms

path: testdata/sha384sum.ll
took: 29.162772ms

path: testdata/sha512sum.ll
took: 28.085743ms

path: testdata/shred.ll
took: 31.52207ms

path: testdata/shuf.ll
took: 36.262054ms

path: testdata/sleep.ll
took: 22.108448ms

path: testdata/sort.ll
took: 83.537074ms

path: testdata/split.ll
took: 25.468708ms

path: testdata/stat.ll
took: 63.992103ms

path: testdata/stdbuf.ll
took: 16.104198ms

path: testdata/stty.ll
took: 30.97863ms

path: testdata/sum.ll
took: 20.282303ms

path: testdata/sync.ll
took: 15.790784ms

path: testdata/tac.ll
took: 18.905382ms

path: testdata/tail.ll
took: 39.719732ms

path: testdata/tee.ll
took: 14.719764ms

path: testdata/test.ll
took: 21.460165ms

path: testdata/timeout.ll
took: 17.223561ms

path: testdata/touch.ll
took: 50.716942ms

path: testdata/tr.ll
took: 24.615551ms

path: testdata/true.ll
took: 12.945751ms

path: testdata/truncate.ll
took: 21.073524ms

path: testdata/tsort.ll
took: 20.569098ms

path: testdata/tty.ll
took: 11.985274ms

path: testdata/uname.ll
took: 18.706545ms

path: testdata/unexpand.ll
took: 16.838912ms

path: testdata/uniq.ll
took: 22.988271ms

path: testdata/unlink.ll
took: 13.343213ms

path: testdata/uptime.ll
took: 29.010324ms

path: testdata/users.ll
took: 22.533472ms

path: testdata/vdir.ll
took: 86.476757ms

path: testdata/wc.ll
took: 19.588786ms

path: testdata/who.ll
took: 17.614309ms

path: testdata/whoami.ll
took: 16.535844ms

path: testdata/yes.ll
took: 13.35127ms

real 3.01
user 2.97
sys 0.04
```
