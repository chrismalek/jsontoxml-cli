package main

import (
	"fmt"
	"os"
)

func main() {

	out := `
 [
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dwo68jPBBSULPMID/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/okta-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/okta-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dwo68jPBBSULPMID/users"
            }
        },
        "id": "00g2dwo68jPBBSULPMID",
        "objectClass": [
            "okta:user_group"
        ],
        "profile": {
            "description": "All users in your organization",
            "name": "Everyone"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrpvHVZXZDHWFE/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrpvHVZXZDHWFE/users"
            }
        },
        "id": "00g2dysrpvHVZXZDHWFE",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/Domain Users",
            "dn": "CN=Domain Users,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "mpDTrW3y9UKepc3xSJn2cg==",
            "groupScope": "Global",
            "groupType": "Security",
            "name": "Domain Users",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-513",
            "samAccountName": "Domain Users",
            "windowsDomainQualifiedName": "NU\\Domain Users"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrqnGNKVENVZED/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrqnGNKVENVZED/users"
            }
        },
        "id": "00g2dysrqnGNKVENVZED",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Builtin/Users",
            "dn": "CN=Users,CN=Builtin,DC=nu,DC=edu",
            "externalId": "JF8b/OfhwEizIy4XyAZa2A==",
            "groupScope": "Domain Local",
            "groupType": "Security",
            "name": "Users",
            "objectSid": "S-1-5-32-545",
            "samAccountName": "Users",
            "windowsDomainQualifiedName": "NU\\Users"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrsrRHRGVPIYVW/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrsrRHRGVPIYVW/users"
            }
        },
        "id": "00g2dysrsrRHRGVPIYVW",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE1_DEPT_TECH",
            "dn": "CN=SHARE1_DEPT_TECH,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "d0byXM7RaUWaP6DJVqYoaw==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE1_DEPT_TECH",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-17283",
            "samAccountName": "SHARE1_DEPT_TECH",
            "windowsDomainQualifiedName": "NU\\SHARE1_DEPT_TECH"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrstTQQTZJNTBL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrstTQQTZJNTBL/users"
            }
        },
        "id": "00g2dysrstTQQTZJNTBL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/VPN",
            "dn": "CN=VPN,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "erZ9hPjUMk63noQpH3godQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "VPN",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-18736",
            "samAccountName": "VPN",
            "windowsDomainQualifiedName": "NU\\VPN"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrt5ACKEEBAHTV/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrt5ACKEEBAHTV/users"
            }
        },
        "id": "00g2dysrt5ACKEEBAHTV",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/ERP - Team Leaders",
            "dn": "CN=ERP - Team Leaders,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "jbvmAU6YAUK4G0CM1kdUMg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "ERP - Team Leaders",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-12525",
            "samAccountName": "ERP - Team Leaders",
            "windowsDomainQualifiedName": "NU\\ERP - Team Leaders"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrunQKWYVSCSIB/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrunQKWYVSCSIB/users"
            }
        },
        "id": "00g2dysrunQKWYVSCSIB",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE22_ERP",
            "dn": "CN=SHARE22_ERP,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "zBtw/G0X7UGdJDdti1eICw==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE22_ERP",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-13736",
            "samAccountName": "SHARE22_ERP",
            "windowsDomainQualifiedName": "NU\\SHARE22_ERP"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrwxFGFKXBCRBS/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrwxFGFKXBCRBS/users"
            }
        },
        "id": "00g2dysrwxFGFKXBCRBS",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/ERP - Team",
            "dn": "CN=ERP - Team,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "Dhlpi/yyw0WubBH1IQgBcw==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "ERP - Team",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-12526",
            "samAccountName": "ERP - Team",
            "windowsDomainQualifiedName": "NU\\ERP - Team"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrwzOPPTMSETWR/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrwzOPPTMSETWR/users"
            }
        },
        "id": "00g2dysrwzOPPTMSETWR",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/DB HRSA",
            "dn": "CN=DB HRSA,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "00F6FDEfD0mAQ6O/S5lIsg==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "DB HRSA",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-13696",
            "samAccountName": "DB HRSA",
            "windowsDomainQualifiedName": "NU\\DB HRSA"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrx1YEYALZUHRU/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrx1YEYALZUHRU/users"
            }
        },
        "id": "00g2dysrx1YEYALZUHRU",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/DB FIN",
            "dn": "CN=DB FIN,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "d/L/t+46+0KiLm4Sq9/1gw==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "DB FIN",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-13693",
            "samAccountName": "DB FIN",
            "windowsDomainQualifiedName": "NU\\DB FIN"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysryzCZSSCOYJRK/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysryzCZSSCOYJRK/users"
            }
        },
        "id": "00g2dysryzCZSSCOYJRK",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/share_doc",
            "dn": "CN=share_doc,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "WaImEZFng0qxIPG9KvwShg==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "share_doc",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-17019",
            "samAccountName": "share_doc",
            "windowsDomainQualifiedName": "NU\\share_doc"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrz1JALUVPOLIL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysrz1JALUVPOLIL/users"
            }
        },
        "id": "00g2dysrz1JALUVPOLIL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/ITG - Info Tech",
            "dn": "CN=ITG - Info Tech,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "iR8twyZnkkuZNtmZJc8zOw==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "ITG - Info Tech",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-11584",
            "samAccountName": "ITG - Info Tech",
            "windowsDomainQualifiedName": "NU\\ITG - Info Tech"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss23RBHKSQOJWV/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss23RBHKSQOJWV/users"
            }
        },
        "id": "00g2dyss23RBHKSQOJWV",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/DBA",
            "dn": "CN=DBA,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "LEoLFEsjAUi5ACO+Um1zdQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "DBA",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-16695",
            "samAccountName": "DBA",
            "windowsDomainQualifiedName": "NU\\DBA"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss3xJHHJVINPRS/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss3xJHHJVINPRS/users"
            }
        },
        "id": "00g2dyss3xJHHJVINPRS",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/dev_pshome",
            "dn": "CN=dev_pshome,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "E2HnhsZvPUmMe5Y/hngJDg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "dev_pshome",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-16781",
            "samAccountName": "dev_pshome",
            "windowsDomainQualifiedName": "NU\\dev_pshome"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss4bUURBIWVAXL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss4bUURBIWVAXL/users"
            }
        },
        "id": "00g2dyss4bUURBIWVAXL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE13_soar",
            "dn": "CN=SHARE13_soar,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "C6FT/RYOCkmq/wuLTqjQqg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE13_soar",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-16733",
            "samAccountName": "SHARE13_soar",
            "windowsDomainQualifiedName": "NU\\SHARE13_soar"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss6bICEIHWVVAJ/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss6bICEIHWVVAJ/users"
            }
        },
        "id": "00g2dyss6bICEIHWVVAJ",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SoftDev",
            "dn": "CN=SoftDev,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "iplc4zcqhEGcX1MXab6XaA==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SoftDev",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-16955",
            "samAccountName": "SoftDev",
            "windowsDomainQualifiedName": "NU\\SoftDev"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss79MPWTITUYZS/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss79MPWTITUYZS/users"
            }
        },
        "id": "00g2dyss79MPWTITUYZS",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/IT Soft Dev",
            "dn": "CN=IT Soft Dev,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "w3FTgp8ZpkGikI8khPWcnA==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "IT Soft Dev",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-12873",
            "samAccountName": "IT Soft Dev",
            "windowsDomainQualifiedName": "NU\\IT Soft Dev"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss9hTDQFBOPFTI/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss9hTDQFBOPFTI/users"
            }
        },
        "id": "00g2dyss9hTDQFBOPFTI",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/ITG - Prog - Analyst",
            "dn": "CN=ITG - Prog - Analyst,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "ZakLRRHV5Ey7KqEDqeH7cQ==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "ITG - Prog - Analyst",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-11659",
            "samAccountName": "ITG - Prog - Analyst",
            "windowsDomainQualifiedName": "NU\\ITG - Prog - Analyst"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss9pYSSUNBJUDJ/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyss9pYSSUNBJUDJ/users"
            }
        },
        "id": "00g2dyss9pYSSUNBJUDJ",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/Programmers",
            "dn": "CN=Programmers,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "fuwnSdTtgEGZMlIwhJ5acg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "Programmers",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-1718",
            "samAccountName": "Programmers",
            "windowsDomainQualifiedName": "NU\\Programmers"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssa7QZFDZAHLQJ/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssa7QZFDZAHLQJ/users"
            }
        },
        "id": "00g2dyssa7QZFDZAHLQJ",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/ITG - Help-Program",
            "dn": "CN=ITG - Help-Program,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "m6FXePsnPkmaCBiXHAimqw==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "ITG - Help-Program",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-11556",
            "samAccountName": "ITG - Help-Program",
            "windowsDomainQualifiedName": "NU\\ITG - Help-Program"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssbhAHDEAKWKJK/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssbhAHDEAKWKJK/users"
            }
        },
        "id": "00g2dyssbhAHDEAKWKJK",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/TALISMA_DCOM_USERS_TLSM-APP",
            "dn": "CN=TALISMA_DCOM_USERS_TLSM-APP,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "A993ozNEUEWWvStUrC7EJg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "TALISMA_DCOM_USERS_TLSM-APP",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-22989",
            "samAccountName": "TALISMA_DCOM_USERS_TLSM-APP",
            "windowsDomainQualifiedName": "NU\\TALISMA_DCOM_USERS_TLSM-APP"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssbnDMMVFTDPVM/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssbnDMMVFTDPVM/users"
            }
        },
        "id": "00g2dyssbnDMMVFTDPVM",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/MIS Programmers",
            "dn": "CN=MIS Programmers,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "EyfzxfFRFkCnYP9jPlR24Q==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "MIS Programmers",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-1865",
            "samAccountName": "MIS Programmers",
            "windowsDomainQualifiedName": "NU\\MIS Programmers"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysscnKUYMASEHZV/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysscnKUYMASEHZV/users"
            }
        },
        "id": "00g2dysscnKUYMASEHZV",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/ERP",
            "dn": "CN=ERP,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "rrQYqLaTXE2NGLm9jVMO5A==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "ERP",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-12598",
            "samAccountName": "ERP",
            "windowsDomainQualifiedName": "NU\\ERP"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssfzZNRWZZLREF/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssfzZNRWZZLREF/users"
            }
        },
        "id": "00g2dyssfzZNRWZZLREF",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/MIS Department",
            "dn": "CN=MIS Department,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "u5OWMUpQUEm9TI1kqJFRNA==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "MIS Department",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-2250",
            "samAccountName": "MIS Department",
            "windowsDomainQualifiedName": "NU\\MIS Department"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssgnYJCYMOGASC/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssgnYJCYMOGASC/users"
            }
        },
        "id": "00g2dyssgnYJCYMOGASC",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/All Users",
            "dn": "CN=All Users,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "Dxt1LjG7ZUSFLaDQq9Ihag==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "All Users",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-11416",
            "samAccountName": "All Users",
            "windowsDomainQualifiedName": "NU\\All Users"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssk3DQSBLEBFHV/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssk3DQSBLEBFHV/users"
            }
        },
        "id": "00g2dyssk3DQSBLEBFHV",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/security_programmers",
            "dn": "CN=security_programmers,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "5glaobUTUEOPfQ8P5rOFLw==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "security_programmers",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-16585",
            "samAccountName": "security_programmers",
            "windowsDomainQualifiedName": "NU\\security_programmers"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssmxHFMGCWPANT/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyssmxHFMGCWPANT/users"
            }
        },
        "id": "00g2dyssmxHFMGCWPANT",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/ITG - Help-All Departments",
            "dn": "CN=ITG - Help-All Departments,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "AFFSUrA6+k2kGTAYTR3zVQ==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "ITG - Help-All Departments",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-11549",
            "samAccountName": "ITG - Help-All",
            "windowsDomainQualifiedName": "NU\\ITG - Help-All"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst4tYBZYKXCPEB/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst4tYBZYKXCPEB/users"
            }
        },
        "id": "00g2dyst4tYBZYKXCPEB",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SoftDevBatch",
            "dn": "CN=SoftDevBatch,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "oum4dwauOEa6iRyEtSnDGw==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SoftDevBatch",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-26874",
            "samAccountName": "SoftDevBatch",
            "windowsDomainQualifiedName": "NU\\SoftDevBatch"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst7lTMDTERQUOL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst7lTMDTERQUOL/users"
            }
        },
        "id": "00g2dyst7lTMDTERQUOL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/WiFi",
            "dn": "CN=WiFi,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "L0yCbHV+F06H1+n9tSSmUA==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "WiFi",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-30316",
            "samAccountName": "WiFi",
            "windowsDomainQualifiedName": "NU\\WiFi"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst7zAJAHRSPUOB/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyst7zAJAHRSPUOB/users"
            }
        },
        "id": "00g2dyst7zAJAHRSPUOB",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/SharePoint Security Groups/VA Readers",
            "dn": "CN=VA Readers,OU=SharePoint Security Groups,DC=nu,DC=edu",
            "externalId": "WyV0DLuK9U6l8iyun5g9pQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "VA Readers",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-29714",
            "samAccountName": "VA Readers",
            "windowsDomainQualifiedName": "NU\\VA Readers"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystdpEZPJFBHRIJ/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystdpEZPJFBHRIJ/users"
            }
        },
        "id": "00g2dystdpEZPJFBHRIJ",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE13_MISPROG$",
            "dn": "CN=SHARE13_MISPROG$,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "r2o1dwLvbkq3YYjjyVa62A==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE13_MISPROG$",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-30872",
            "samAccountName": "SHARE13_MISPROG$",
            "windowsDomainQualifiedName": "NU\\SHARE13_MISPROG$"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystdrTYVJVPHAUU/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystdrTYVJVPHAUU/users"
            }
        },
        "id": "00g2dystdrTYVJVPHAUU",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE13_MISADMIN",
            "dn": "CN=SHARE13_MISADMIN,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "uK7sqaB0OkCy900J4E2QAQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE13_MISADMIN",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-30873",
            "samAccountName": "SHARE13_MISADMIN",
            "windowsDomainQualifiedName": "NU\\SHARE13_MISADMIN"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyste7AWWGWDCSDO/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dyste7AWWGWDCSDO/users"
            }
        },
        "id": "00g2dyste7AWWGWDCSDO",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE98_PStraining",
            "dn": "CN=SHARE98_PStraining,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "/mmx6oaI2ka/uDVGFhS4cg==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE98_PStraining",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-30884",
            "samAccountName": "SHARE98_PStraining",
            "windowsDomainQualifiedName": "NU\\SHARE98_PStraining"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystebAYWVZDQGUM/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystebAYWVZDQGUM/users"
            }
        },
        "id": "00g2dystebAYWVZDQGUM",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/SHARE54_CRM_ATTACHMENTS_WRITE",
            "dn": "CN=SHARE54_CRM_ATTACHMENTS_WRITE,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "rHkHd2cdXEenxpoDjFM6/A==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SHARE54_CRM_ATTACHMENTS_WRITE",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-30887",
            "samAccountName": "SHARE54_CRM_ATTACHMENTS_WRITE",
            "windowsDomainQualifiedName": "NU\\SHARE54_CRM_ATTACHMENTS_WRITE"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystffIEKDCDGZXL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystffIEKDCDGZXL/users"
            }
        },
        "id": "00g2dystffIEKDCDGZXL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/upgrade",
            "dn": "CN=upgrade,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "mT6IWcbK+0qI0dyfpFV+VQ==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "upgrade",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-31573",
            "samAccountName": "upgrade",
            "windowsDomainQualifiedName": "NU\\upgrade"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysthjZZCKVURMCB/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysthjZZCKVURMCB/users"
            }
        },
        "id": "00g2dysthjZZCKVURMCB",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Security Groups/ISIS_Share",
            "dn": "CN=ISIS_Share,OU=Security Groups,DC=nu,DC=edu",
            "externalId": "6Z0I5jOfxE67b9mAlP9+UQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "ISIS_Share",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-33768",
            "samAccountName": "ISIS_Share",
            "windowsDomainQualifiedName": "NU\\ISIS_Share"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysthnIEPIAIEJNL/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dysthnIEPIAIEJNL/users"
            }
        },
        "id": "00g2dysthnIEPIAIEJNL",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/NUHRUPG",
            "dn": "CN=NUHRUPG,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "hW0dez8T402RWdS9SLNo+Q==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "NUHRUPG",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-34225",
            "samAccountName": "NUHRUPG",
            "windowsDomainQualifiedName": "NU\\NUHRUPG"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystnrRFAHNBISMR/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2dystnrRFAHNBISMR/users"
            }
        },
        "id": "00g2dystnrRFAHNBISMR",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/CityU Project",
            "dn": "CN=CityU Project,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "vdu5IEz2V0S7NQSie4gl4Q==",
            "groupScope": "Universal",
            "groupType": "Distribution",
            "name": "CityU Project",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-38262",
            "samAccountName": "CityU Project",
            "windowsDomainQualifiedName": "NU\\CityU Project"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2k1242zYVLXYXGYTI/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2k1242zYVLXYXGYTI/users"
            }
        },
        "id": "00g2k1242zYVLXYXGYTI",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/Distribution Lists/NU-SSO",
            "dn": "CN=NU-SSO,OU=Distribution Lists,DC=nu,DC=edu",
            "externalId": "BRNHHNkTl0y2A+aGLShA2Q==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "NU-SSO",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-45980",
            "samAccountName": "NU-SSO",
            "windowsDomainQualifiedName": "NU\\NU-SSO"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2k5dohfTGFMUUPEJJ/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/active_directory-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2k5dohfTGFMUUPEJJ/users"
            }
        },
        "id": "00g2k5dohfTGFMUUPEJJ",
        "objectClass": [
            "okta:windows_security_principal"
        ],
        "profile": {
            "description": "nu.edu/SharePoint Security Groups/SSO Contributors",
            "dn": "CN=SSO Contributors,OU=SharePoint Security Groups,DC=nu,DC=edu",
            "externalId": "iPeFv9c0ok24QRNeglLyIQ==",
            "groupScope": "Universal",
            "groupType": "Security",
            "name": "SSO Contributors",
            "objectSid": "S-1-5-21-82779875-1711884744-1244796221-50178",
            "samAccountName": "SSO Contributors",
            "windowsDomainQualifiedName": "NU\\SSO Contributors"
        }
    },
    {
        "_links": {
            "apps": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2em1dp5MILVBUXNMI/apps"
            },
            "logo": [
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/okta-medium.png",
                    "name": "medium",
                    "type": "image/png"
                },
                {
                    "href": "https://nu.oktapreview.com/img/logos/groups/okta-large.png",
                    "name": "large",
                    "type": "image/png"
                }
            ],
            "users": {
                "href": "https://nu.oktapreview.com/api/v1/groups/00g2em1dp5MILVBUXNMI/users"
            }
        },
        "id": "00g2em1dp5MILVBUXNMI",
        "objectClass": [
            "okta:user_group"
        ],
        "profile": {
            "description": "Access to PeopleSoft Dev Sites",
            "name": "peopleSoftDevAccess"
        }
    }
]
`

	fmt.Print(`{"root_element":{"foo":"bar1"}}`)
	// fmt.Print(out)
	if len(out) < 2 {

	}
	// fmt.Print(`shit`)
	os.Exit(0)
}
