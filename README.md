Downloads IEEE OUI, OUI28 and OUI36 data from official site.

Saves OUI data in "oui" hash.

Company names only, no addresses.

MAC search order:

lookup "oui" hash with first 6 half-octets (3 octets)

If you find "IEEE Registration Authority", then lookup "oui" with first 7 or 9 half-octets.



Hash keys are lowercase!

Example:

```
$ redis-cli HGETALL oui | grep -A 1 "^d01411" | head
d01411
IEEE Registration Authority
--
d01411b
CYLTek Limited
--
d014115
Superlead
--
d014119

```

```
$ redis-cli HGET oui 70b3d5
IEEE Registration Authority

$ redis-cli HGETALL oui | grep -A 1 70b3d5 | head
70b3d56c0
LLC "NTZ "Mekhanotronika"
--
70b3d5978
Satixfy Israel Ltd.
--
70b3d5968
LGM Ingénierie
--
```
