#!/usr/bin/env python3

import glob
import subprocess
import os


def can_exec(name):
    for path in os.environ['PATH'].split(os.pathsep):
            exe_file = os.path.join(path, name)
            if os.path.isfile(exe_file) and os.access(exe_file, os.X_OK):
                return exe_file
    return False


def check_system():
    if can_exec('swagger-22') == False:
        raise EnvironmentError('Unable to locate swagger-22 in $PATH')


def get_targets():
    """Gets a list of targets by scanning the current directory for
    Swagger files and then returns a list of tuples being (name, source)
    """
    out = list()
    for fileName in glob.glob('*.swagger.json'):
        name = fileName[:fileName.index('.')]
        out.append((name, fileName))
    return out


def prepare_target(name):
    subprocess.call(['rm', '-rf', name])



def build_target(name, src):
    subprocess.call(['rm', '-rf', name])
    subprocess.call(['mkdir', '-p', name])
    subprocess.call([
        'swagger-22', 'generate', 'client',
        '-f', src,
        '-A', name,
        '-t', name,
    ])


def main():
    for target in get_targets():
        prepare_target(target[0])
        build_target(target[0], target[1])


if __name__ == '__main__':
    check_system()
    main()
