# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.


# UbuntuSystemProcesses contain a list of Ubuntu specific system
# processes.
UbuntuSystemProcesses = [
    "acpi",
    "ata",
    "bioset",
    "charger",
    "crypto",
    "deferwq",
    "devfreq",
    "ecryptfs",
    "ext4",
    "fsnotify",
    "ib",
    "ipv6",
    "iscsi",
    "iw",
    "ixgbevf",
    "jbd2",
    "jfsCommit",
    "jfsIO",
    "jfsSync",
    "kauditd",
    "kblockd",
    "kdevtmpfs",
    "khugepaged",
    "khungtaskd",
    "kintegrityd",
    "kpsmoused",
    "ksmd",
    "ksoftirqd",
    "kswapd0",
    "kthreadd",
    "kthrotld",
    "kworker",
    "/lib/systemd/systemd",
    "/lib/systemd/systemd-journald",
    "/lib/systemd/systemd-logind",
    "/lib/systemd/systemd-timesyncd",
    "/lib/systemd/systemd-udevd",
    "md",
    "migration",
    "netns",
    "perf",
    "raid5wq",
    "rcu",
    "rdma",
    "/sbin/agetty",
    "/sbin/dhclient (deleted)",
    "/sbin/iscsid",
    "/sbin/lvmetad",
    "/sbin/mdadm",
    "scsi",
    "ttm",
    "/usr/bin/dbus-daemon",
    "/usr/bin/docker-containerd",
    "/usr/bin/dockerd",
    "/usr/bin/lxcfs (deleted)",
    "/usr/lib/accountsservice/accounts-daemon",
    "/usr/lib/policykit-1/polkitd",
    "/usr/lib/snapd/snapd",
    "/usr/sbin/acpid",
    "/usr/sbin/atd",
    "/usr/sbin/cron",
    "/usr/sbin/dnsmasq",
    "/usr/sbin/irqbalance",
    "/usr/sbin/rsyslogd",
    "/usr/sbin/sshd",
    "vmstat",
    "watchdog",
    "writeback",
    "xenbus",
    "xenwatch",
    "xfs",
    "xfsalloc",


]

# SystemProcesses contains a list of the system processes that are
# installed by the system by default.

SystemProcesses = UbuntuSystemProcess