Downloads IEEE OUI, OUI28 and OUI36 data from official site.

Saves 24-bit OUI data in "oui" hash.

Saves 28-bit OUI data in "oui28" hash.

Saves 36-bit OUI data in "oui36" hash.

Company names only, no addresses.

MAC search order:

lookup "oui" hash with first 6 half-octets (3 octets)

If you find "IEEE Registration Authority", then lookup "oui28" with first 7 half-octets or lookup "oui36" with first 9 half-octets until you find answer (if it is there).



Hash keys are lowercase!

Example:

```
$ redis-cli HGETALL oui | head
28bb59
RNET Technologies, Inc.
a8610a
ARDUINO AG
c853e1
Beijing Bytedance Network Technology Co., Ltd
a06610
FUJITSU LIMITED
a42249
Sagemcom Broadband SAS

$ redis-cli HGETALL oui28 | head
c0eac3d
Kontron Asia Technology Inc.
4ce1734
Huizhou Dehong Technology Co., Ltd.
80a5791
Zhe Jiang EV-Tech Co.,Ltd
c0482fc
Lunar USA Inc.
98aafc9
BEAM Authentic

$ redis-cli HGETALL oui36 | head
70b3d594d
SEASON DESIGN TECHNOLOGY
70b3d50c2
LOOK EASY INTERNATIONAL LIMITED
70b3d5f9a
Krabbenhøft og Ingolfsson
70b3d5abe
MART NETWORK SOLUTIONS LTD
70b3d59e8
Zerospace ICT Services B.V.
```
