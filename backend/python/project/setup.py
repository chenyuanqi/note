# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


with open('README.rst') as f:
    readme = f.read()

with open('LICENSE') as f:
    license = f.read()

setup(
    name='project',
    version='0.0.1',
    description='python project',
    long_description=readme,
    author='Kenneth Reitz',
    author_email='chenyuanqi@outlook.com',
    url='https://github.com/chenyuanqi/python-project',
    license=license,
    packages=find_packages(exclude=('tests',))
)

