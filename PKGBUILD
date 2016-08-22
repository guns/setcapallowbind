# Maintainer: Sung Pae <self@sungpae.com>
pkgname=setcapallowbind
pkgver=0
pkgrel=1
pkgdesc='Single static binary that calls `setcap cap_net_bind_service=+ep`'
license=('MIT')
arch=('x86_64')
groups=('nerv')

pkgver() {
    git describe --long --tags | tr - .
}

package() {
    cd "$startdir"
    mkdir -p "pkg/$pkgname/usr/bin"
    go build -i -v -o "pkg/$pkgname/usr/bin/setcapallowbind"
}
