// Copyright © 2019 Ettore Di Giacinto <mudler@gentoo.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package helpers

import (
	"io"
	"os"

	"github.com/docker/docker/pkg/archive"
)

func Tar(src, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	fs, err := archive.Tar(src, archive.Uncompressed)
	if err != nil {
		return err
	}
	defer fs.Close()

	_, err = io.Copy(out, fs)
	if err != nil {
		return err
	}

	err = out.Sync()
	if err != nil {
		return err
	}
	return err
}

// Untar just a wrapper around the docker functions
func Untar(src, dest string, sameOwner bool) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	return archive.Untar(in, dest, &archive.TarOptions{
		NoLchown:        !sameOwner,
		ExcludePatterns: []string{"dev/"}, // prevent 'operation not permitted'
	})
}