# nsec3label
Compute nsec3 hash for a given label.

## Background

Many TLD's and domain owners sign their zones with DNSSEC.
To mitigate the risk of zone enumeration through the NSEC chain, [RFC5155](https://tools.ietf.org/html/rfc5155) introduced NSEC3.

NSEC3 records work by first hashing all labels left of the zone name.

Here is how an NSEC record looks like
```
iis.se.			7200	IN	NSEC	iis07.se. NS DS RRSIG NSEC
```

Here is how an NSEC3 record looks like
```
uef535kvt93ligh45oa3qt4c0gpmgv55.nu. 7200 IN NSEC3 1 0 5 4BC4A733E155C9A1 UEF6NV3QN3JCOALUD1SQO8Q2IVNN64JS NS DS RRSIG
```

When there is an error in the NSEC3 chain of your zone you might want to check some of your NSEC3 records. 
In that case it would be very helpful to know which domainname is hashed to which label.

nsec3label solves this problem.

## Usage
Input is all information needed to hash create the hash, namely
- hash algorithm number (1=SHA1)
- iterations (number of hash iterations)
- salt 
- fqdn fully qualified domain name

These parameters are easiest to get from the NSEC3PARAM resource record.

Example from .nu
```
nu.			1961	IN	NSEC3PARAM 1 0 5 4BC4A733E155C9A1
```

How to invoke the script
```
nsec3label nsec3label 1 5 4BC4A733E155C9A1 iis.nu
```

Output
```
UEF535KVT93LIGH45OA3QT4C0GPMGV55
```

