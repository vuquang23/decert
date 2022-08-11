#!/bin/bash
#
# dnslcheck.sh
#
# Last edited: 11/22/2017
#
# Copyright (c) 2017 by Catherine A. Jefferson <ariel@spambouncer.org>
#
#   Released under the provisions of the GNU General Public License,
#   version 3 (https://www.gnu.org/copyleft/gpl.txt).
#
#   Everyone is permitted to copy and distribute this script, and
#   modify it as they wish, but must make any modified versions
#   available under the same provisions as this version.
#
# Description:
#  Takes an IP address, domain or email address, and tells you if
#  it is on the specified DNS-based list or appropriate class of lists.
#  If no list is specified, the default list for that data class is
#  checked. If the keyword "all" is specified, all supported DNS-based
#  lists for that data type are checked.
#
# Supported types of list:
#  * IP-based (DNSBL): IP address
#  * domain-based (RHSBL or URIBL): domain
#  * email-based (HASHBL): email address
#
# Syntax:
#  dnslcheck.sh <(IP|domain|email)> [ <list] ... ]

## Get the command line parameters.

if [ "${1}" == "help" ] || [ "${1}" == "h" ] || [ "${1}" == "?" ]; then
 echo 'Specify an IP address, domain, or hash, and optional dns-based'
 echo 'list or tag for a group of lists.'
 echo ' '
 echo 'Syntax: dnslcheck.sh <IP|domain|hash> [ <list] ... ]'
 echo ' '
 echo '  Looking up an IP:'
 echo '    dnslcheck.sh 127.0\.0\.2      # Look up IP on default IP BLs.'
 echo '    dnslcheck.sh 127.0\.0\.2 sbl  # Look up IP on SBL.'
 echo '    dnslcheck.sh 127.0\.0\.2 all  # Look up IP on all IP BLs.'
 echo ' '
 echo '  Looking up a domain:'
 echo '    dnslcheck.sh example.com      # Look up domain on default domain BLs.'
 echo '    dnslcheck.sh example.com dbl  # Look up domain on DBL.'
 echo '    dnslcheck.sh example.com all  # Look up domain on all domain BLs.'
 echo ' '
 echo '  Looking up an email address:'
 echo '    dnslcheck.sh <noemail@example.com>    # Look up email address on EBL.'
 echo ' '
 echo 'For a list of valid blocklists, type "dnslcheck.sh l" or'
 echo '"dnslcheck.sh list".'
 exit 1
elif [ "${1}" == "list" ] || [ "${1}" == "l" ]; then
 echo 'IP lists and tags:'
 echo '  default (default): Looks up Spamhaus IP-based lists & Spamcop'
 echo '  all: Looks up *every* IP list. (Can be slow)'
 echo '  sbl: Spamhaus SBL (http://www.spamhaus.org/sbl)'
 echo '  drop: Spamhaus DROP/EDrOP (https://www.spamhaus.org/drop)'
 echo '  css: Spamhaus CSS (http://www.spamhaus.org/css)'
 echo '  xbl: Spamhaus XBL (http://www.spamhaus.org/xbl)'
 echo '  pbl: Spamhaus PBL (http://www.spamhaus.org/pbl)'
 echo '  spamcop OR scbl: Spamcop BL'
 echo '     (http://www.spamcop.net/bl.shtml)'
 echo '  barracuda OR brbl: Barracuda BL'
 echo '     (http://barracudacentral.org/rbl)'
 echo '  cbl: Composite BL (http://cbl.abuseat.org)'
 echo '  cymru: Cymru Bogons BL'
 echo '     (http://www.team-cymru.org/Services/Bogons/)'
 echo '  cymru.full: Cymru Full IPv4 Bogons BL'
 echo '     (http://www.team-cymru.org/Services/Bogons/)'
 echo '  dronebl: DroneBL.org BL (http://dronebl.org/)'
 echo '  fabel.dk: Fabel.DK Spam Sources BL'
 echo '     (http://www.fabel.dk/relay/)'
 echo '  interserver: Interserver BL (http://rbl.interserver.net/)'
 echo '  mailspike: Mailspike Reputation List (bad only)'
 echo '      (http://mailspike.net/about.html)'
 echo '  psbl: Passive Spam BL (http://psbl.org)'
 echo '  rbl.jp: RBL.jp BL (http://www.rbl.jp/)'
 echo '  rpss: ReturnPath SenderScore Reputation List (http://senderscore.com/)'
 echo '  scispam: scientificspam.net BL (http://scientificspam.net/)'
 echo '  sorbs: SORBS BL (http://sorbs.net/)'
 echo '  suomispam: suomispam.net BL (http://suomispam.net/)'
 echo '  suretymail: ISIPP SuretyMail WL (http://isipp.com/)'
 echo '  swinog: Swinog IP BL (ImproWare AG) (http://antispam.imp.ch/)'
 echo '  wpbl: Weighted Private BL (http://www.wpbl.info/)'
 echo 'Domain lists and tags:'
 echo '  default (default): Looks up Spamhaus DBL & SURBL'
 echo '  all: Looks up *every* domain list. (Can be slow)'
 echo '  dbl: Spamhaus Domain Blocklist (http://spamhaus.org/dbl/)'
 echo '  suomispam: suomispam.net domain BL (http://suomispam.net/)'
 echo '  sarbl: Smart Adaptative Realtime URIBL (http://sarbl.org/)'
 echo '  surbl: Spam URI Blocklist (http://surbl.org/)'
 echo '  uribl: URI Blocklist (http://uribl.com/)'
 echo '  scispam: scientificspam.net BL (http://scientificspam.net/)'
 echo '  swinogdom: Swinog Domain BL (ImproWare AG) (http://antispam.imp.ch/)'
 echo 'Email address lists and tags:'
 echo '  default (default): MSBL EBL (http://msbl.org/ebl.html). '
 echo '  ("all" and "ebl" do the same thing.)'
 exit 1
elif [[ "${1}" =~ ^[0-9][0-9]?[0-9]?\.[0-9][0-9]?[0-9]?\.[0-9][0-9]?[0-9]?\.[0-9][0-9]?[0-9]?$ ]]; then
 type="IP"
 IP=${1}
elif [[ "${1}" =~ ^(test|([0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]*\.)*([Xx][Nn]--[0-9A-Za-z][0-9A-Za-z]*|([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)*))$ ]]; then
 type="domain"
 domain=${1}
elif [[ "${1}" =~ ^[0-9A-Za-z][0-9A-Za-z._$+-]*@([0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]*\.)*([Xx][Nn]--[0-9A-Za-z][0-9A-Za-z]*|([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)*)$ ]]; then
 type="hash"
 hashchk=${1}
elif [[ "${1}" =~ ^hash:.*$ ]]; then
 type="hash"
 hashchk=${1}
else
 echo ' Invalid input! Please specify a valid IPv4 IP address, domain name,'
 echo ' or SHA1 hash.'
 exit 1
fi

if [ "${type}" == "hash" ]; then
  if [ -e "/usr/bin/sha1sum" ]; then
    shasum='/usr/bin/sha1sum'
  elif [ -e "/bin/sha1" ]; then
    shasum='/bin/sha1'
  else
    echo 'WARNING! SHA1 creation utility not found. Cannot look up '
    echo '  email address!'
    exit 1
  fi
fi

## Set the list of blocklists to be checked.

if [ "${type}" == "IP" ]; then
 if [[ ${2} =~ ^(default|all|sbl|drop|css|xbl|pbl|spamcop|scbl|barracuda|brbl|cbl|cymru|cymru.full|drbl|dronebl|fabel.dk|interserver|iadb|isipp|mailspike|psbl|rbl.jp|returnpath|rpss|scispam|senderscore|sorbs|suomispam|suretymail|swinog|wpbl)$ ]]; then
  LIST=${2}
 elif [ "${2}" == "" ]; then
  LIST="default"
 else
  echo ' Invalid blocklist name or tag for IP lookups!  Please specify'
  echo ' a valid IP blocklist name, "all", "default", or leave blank to'
  echo ' check the default IP blocklists.'
  exit 1
 fi
elif [ "${type}" == "domain" ]; then
 if [[ ${2} =~ ^(default|all|dbl|sarbl|scispam|suomispam|surbl|uribl|swinogdom)$ ]]; then
  LIST=${2}
 elif [ "${2}" == "" ]; then
  LIST="default"
 else
  echo ' Invalid blocklist name or tag for domain lookups!  Please specify'
  echo ' a valid domain blocklist name, "all", "default", or leave blank to'
  echo ' check the default domain blocklists.'
  exit 1
 fi
elif [ "${type}" == "hash" ]; then
 if [[ ${2} =~ ^(default|all|dropbox|ebl)$ ]]; then
  LIST=${2}
 elif [ "${2}" == "" ]; then
  LIST="default"
 else
  echo ' Invalid blocklist name or tag for email lookups!  Please specify'
  echo ' a valid email blocklist name, "all", "default", or leave blank to'
  echo ' check the default email blocklists.'
  exit 1
 fi
else
 echo ' Something unexpected happened. Please report this to'
 echo ' <ariel@spambouncer.org>.'
 exit 1
fi

## Prepare the input data
##
## (Reverse IP addresees, lowercase domains, canonicalize and
## hash hash: strings.)

if [ "${type}" == "IP" ]; then
 D1='NULL'
 D2='NULL'
 D3='NULL'
 L1='NULL'
 L2='NULL'
 L3='NULL'
 L4='NULL'
 L4a='NULL'
 LB='NULL'
 Q1='NULL'
 Q2='NULL'
 Q3='NULL'
 Q4='NULL'
 REVIP='NULL'
 D1=`expr "$IP" : '[0-9][0-9]*[.]'`
 D2=`expr "$IP" : '[0-9][0-9]*[.][0-9][0-9]*[.]'`
 D3=`expr "$IP" : '[0-9][0-9]*[.][0-9][0-9]*[.][0-9][0-9]*[.]'`
 L1=`expr $D1 - 1`
 L2=`expr $D2 - $D1 - 1`
 L3=`expr $D3 - $D2 - 1`
 L4a=`echo ${#IP}`
 L4=`expr ${L4a} - ${D3} + 1`
 Q1=${IP:0:$L1}
 Q2=${IP:$D1:$L2}
 Q3=${IP:$D2:$L3}
 Q4=${IP:$D3:$L4}
 REVIP="${Q4}.${Q3}.${Q2}.${Q1}"
elif [ "${type}" == "domain" ]; then
 domain=`echo ${domain} | tr '[:upper:]' '[:lower:]'`
elif [ "${type}" == "hash" ]; then
 hashchk=`echo ${hashchk} | sed -e 's/^hash://g'`
 if [[ "${hashchk}" =~ ^[0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_+.0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]*@([0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]*\.)+([Xx][Nn]--[0-9A-Za-z][0-9A-Za-z]*|([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)*)$ ]]; then
  hashstr=`echo ${hashchk} | sed -e 's/^\([0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_.0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]*\)[+][^@]*@\([0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ][-_0-9A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]\.([Xx][Nn]--[0-9A-Za-z][0-9A-Za-z]*|([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)([A-ZŞßÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏĞÑÒÓÔÕÖØŒŠÙÚÛÜİŞŸa-zŞßàáâãäåæçèéêëìíîïğñòóôõöøœšùúûüışÿ]|\?)*)\)$/\1@\2/g'`
  hashstr=`echo ${hashstr} | tr '[:upper:]' '[:lower:]'`
  if [[ ${hashstr} =~ @gmail\.com$ ]]; then
   hstr1=`echo ${hashstr} | sed -e 's/^\(.*\)@.*$/\1/g'`
   hstr2=`echo ${hashstr} | sed -e 's/^.*@\(.*\)$/\1/g'`
   hstr1=`echo ${hstr1} | tr -d '.'`
   hashstr="${hstr1}@${hstr2}"
  fi
 fi
 hash=`echo -n ${hashstr} | ${shasum}`
 hash=`echo -n ${hash} | sed -e 's/[^0-9A-Fa-f]*$//g'`
fi

## Do the lookup(s).

### IP Blocklists

if [ "${type}" == "IP" ]; then

 if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "sbl" ] || [ "${LIST}" == "drop" ] || [ "${LIST}" == "css" ] || [ "${LIST}" == "xbl" ] || [ "${LIST}" == "pbl" ] || [ "${LIST}" == "zen" ]; then
   LB=`dig +short ${REVIP}.zen.spamhaus.org 2> /dev/null`

   if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "sbl" ] || [ "${LIST}" == "zen" ]; then
     if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
       LB2='NULL'
       LB2=`dig +short -t txt ${REVIP}.sbl.spamhaus.org 2> /dev/null`
       LB2=`echo ${LB2} | sed -e 's#\"https://www\.spamhaus\.org\/sbl\/query\/SBLCSS\"##g'`
       LB2=`echo ${LB2} | sed -e 's#\"https://www\.spamhaus\.org\/sbl\/query\/\(SBL[0-9][0-9]*\)\"#\1#g'`
       echo "${IP} is in the SBL (${LB2})"
     fi
   fi

   if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "drop" ] || [ "${LIST}" == "zen" ]; then
     if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.9([^0-9.]|$) ]]; then
       LB2='NULL'
       LB2=`dig +short -t txt ${REVIP}.sbl.spamhaus.org 2> /dev/null`
       LB2=`echo ${LB2} | sed -e 's#\"https://www\.spamhaus\.org\/sbl\/query\/SBLCSS\"##g'`
       LB2=`echo ${LB2} | sed -e 's#\"https://www\.spamhaus\.org\/sbl\/query\/\(SBL[0-9][0-9]*\)\"#\1#g'`
       echo "${IP} is in DROP/EDROP (${LB2})"
     fi
   fi

   if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "css" ] || [ "${LIST}" == "zen" ]; then
     if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.3([^0-9.]|$) ]]; then
       echo "${IP} is in the CSS"
     fi
   fi

   if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "xbl" ] || [ "${LIST}" == "zen$ ]; then
     if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.4([^0-9.]|$) ]]; then
       echo "${IP} is in the XBL"
     fi
   fi

   if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "pbl" ] || [ "${LIST}" == "zen$ ]; then
     if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.10([^0-9.]|$) ]] || [[ ${LB} =~ 127\.0\.0\.11 ]]; then
       echo "${IP} is in the PBL"
     fi
   fi

 fi

 if [ "${LIST}" == "default" ] || [ "${LIST}" == "all" ] || [ "${LIST}" == "spamcop" ] || [ "${LIST}" == "scbl" ]; then
   LB=`dig +short ${REVIP}.bl.spamcop.net 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in Spamcop"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "brbl" ] || [ "${LIST}" == "barracuda" ]; then
   LB=`dig +short ${REVIP}.b.barracudacentral.org 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in the BRBL"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "cbl" ]; then
   LB=`dig +short ${REVIP}.cbl.abuseat.org 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in the CBL"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "cymru" ]; then
   LB=`dig +short ${REVIP}.bogons.cymru.com 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in Cymru (Bogons)"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "cymru.full" ]; then
   LB=`dig +short ${REVIP}.v4.fullbogons.cymru.com 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in Cymru (IPv4 Full Bogons)"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "dronebl" ] || [ "${LIST}" == "drbl" ]; then
   LB=`dig +short ${REVIP}.dnsbl.dronebl.org 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.3([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (IRC Drone)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.5([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Bottler)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.6([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Unknown Bot)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.7([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (DDoS Drone)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.8([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (SOCKS Proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.9([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (HTTP Proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.10([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Proxy Chain)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.11([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Web Page Proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.12([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Open DNS Resolver)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.13([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Brute Force Attacker)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.14([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Open Wingate)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.15([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Compromised Router)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.16([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Autorooting Worm)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.17([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Autodetected Botnet IP)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.18([^0-9.]|$) ]]; then
     echo "${IP} is in DroneBL (Non-IRC Hostname)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "fabel.dk" ]; then
   LB=`dig +short ${REVIP}.spamsources.fabel.dk 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in Fabel.dk"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "interserver" ]; then
   LB=`dig +short ${REVIP}.rbl.interserver.net 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in Interserver"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "mailspike" ]; then
   LB=`dig +short ${REVIP}.rep.mailspike.net 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.10([^0-9.]|$) ]]; then
     echo "${IP} is in Mailspike (L5-Worst)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.11([^0-9.]|$) ]]; then
     echo "${IP} is in Mailspike (L4-Very Bad)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.12([^0-9.]|$) ]]; then
     echo "${IP} is in Mailspike (L3-Bad)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.13([^0-9.]|$) ]]; then
     echo "${IP} is in Mailspike (L2-Suspicious)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "psbl" ]; then
   LB=`dig +short ${REVIP}.psbl.surriel.com 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in the PSBL"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "rbl.jp" ]; then
   LB=`dig +short ${REVIP}.all.rbl.jp 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in the RBL.jp (Virus)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.4([^0-9.]|$) ]]; then
     echo "${IP} is in the RBL.jp (Dynamic)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "returnpath" ] || [ "${LIST}" == "rpss" ] || [ "${LIST}" == "senderscore" ]; then
   LB=`dig +short ${REVIP}.score.senderscore.com 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.4\.[0-9][0-9]?[0-9]?([^0-9.]|$) ]]; then
     score=`echo -n ${LB} | sed -e 's/^127\.0\.4\.\(.*\)$/\1/g'`
     echo "${IP} has a SenderScore of ${score}"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "scispam" ]; then
   LB=`dig +short ${REVIP}.bl.scientificspam.net 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     LB2=`dig txt +short ${REVIP}.bl.scientificspam.net 2> /dev/null`
     echo "${IP} is in ScientificSpam (Reason: ${LB2}"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "sorbs" ]; then
   LB=`dig +short ${REVIP}.dnsbl.sorbs.net 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (http proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.3([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (socks proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.4([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (other proxy)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.5([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (open relay)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.6([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (direct spam)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.7([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (vulnerble website)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.8([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (do not test)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.9([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (zombie)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.10([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (dynamic IP)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.11([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (badconf)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.12([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (nomail"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.14([^0-9.]|$) ]]; then
     echo "${IP} is in SORBS (noserver)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "suomispam" ]; then
   LB=`dig +short ${REVIP}.bl.suomispam.net 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in SuomiSpam (black)"
   fi

   LB=`dig +short ${REVIP}.gl.suomispam.net 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in SuomiSpam (grey)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "iadb" ] || [ "${LIST}" == "isipp" ] || [ "${LIST}" == "suretymail" ]; then
   LB=`dig +short ${REVIP}.iadb.isipp.com 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.255([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail (vouched)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.3\.100\.[7-9]([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail (Bulk SOI)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.3\.100\.(10|1100)([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail (Bulk COI)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.3\.100\.200([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail (Transactional)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.3\.100\.211([^0-9.]|$) ]]; then
     echo "${IP} is in SuretyMail (Social Networking)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "swinog" ]; then
   LB=`dig +short ${REVIP}.dnsrbl.swinog.ch 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.3([^0-9.]|$) ]]; then
     echo "${IP} is in Swinog"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "wpbl" ]; then
   LB=`dig +short ${REVIP}.db.wpbl.info 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${IP} is in the WPBL"
   fi
 fi

fi

if [ "${type}" == "domain" ]; then

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "default" ] || [ "${LIST}" == "dbl" ]; then
   LB=`dig +short ${domain}.dbl.spamhaus.org 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.2([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (spam)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.4([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (phish)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.5([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (malware)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.6([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (botnet C&C)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.102([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit spam)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.103([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit redirector)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.104([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit phish)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.105([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit malware)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.106([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit spam)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.102([^0-9.]|$) ]]; then
     echo "${domain} is in the DBL (abused legit botnet C&C)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "default" ] || [ "${LIST}" == "surbl" ]; then
   LB=`dig +short ${domain}.multi.surbl.org 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.(64|72|80|88|192|200|208|216)([^0-9.]|$) ]]; then
     echo "${domain} is in SURBL (Abuse)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.(8|24|72|88|136|144|152|216)([^0-9.]|$) ]]; then
     echo "${domain} is in SURBL (Phish)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.(16|24|80|144|152|208|216)([^0-9.]|$) ]]; then
     echo "${domain} is in SURBL (Malware)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.(128|136|144|152|216)([^0-9.]|$) ]]; then
     echo "${domain} is in SURBL (Legit/Cracked)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "uribl" ]; then
   LB=`dig +short ${domain}.multi.uribl.com 2> /dev/null`

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${domain} is in URIBL (black)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.4([^0-9.]|$) ]]; then
     echo "${domain} is in URIBL (grey)"
   fi

   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.8([^0-9.]|$) ]]; then
     echo "${domain} is in URIBL (red)"
   fi

 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "sarbl" ]; then
   LB=`dig +short ${domain}.public.sarbl.org 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${domain} is in SARBL"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "scispam" ]; then
   LB=`dig +short ${domain}.rhsbl.scientificspam.net 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.2([^0-9.]|$) ]]; then
     LB2=`dig txt +short ${domain}.rhsbl.scientificspam.net 2> /dev/null`
     LB2=`echo ${LB2} | sed -e 's#^\"\(.*\)\"$#\1#g'`
     echo "${domain} is in ScientificSpam (${LB2})"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "suomispam" ]; then
   LB=`dig +short ${domain}.dbl.suomispam.net 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.1\.2([^0-9.]|$) ]]; then
     echo "${domain} is in Suomispam"
   fi
 fi

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "swinogdom" ]; then
   LB=`dig +short ${domain}.uribl.swinog.ch 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     echo "${domain} is in SARBL"
   fi
 fi

fi

if [ "${type}" == "hash" ]; then

 if [ "${LIST}" == "all" ] || [ "${LIST}" == "default" ] || [ "${LIST}" == "dropbox" ] || [ "${LIST}" == "ebl" ]; then
   LB=`dig +short -t A ${hash}.ebl.msbl.org 2> /dev/null`
   if [[ ${LB} =~ (^|[^0-9.])127\.0\.0\.2([^0-9.]|$) ]]; then
     LB2=`dig +short -t txt ${hash}.ebl.msbl.org 2> /dev/null`
     LB2=`echo ${LB2} | sed -e 's#^\"\(.*\)\"$#\1#g'`
     echo "${hashchk} is in the EBL (${LB2})"
   fi
 fi

fi
