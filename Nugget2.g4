grammar Nugget2;

prog: (nug | NL
        )*
        EOF
;


nug:   ID ntype '[]'?
;

ntype: 'string' |
      'sha1hash' |
      'md5hash'  |
      'ntfs'     |
      'file'     |
      'packet'   |
      'exifinfo' |
;

INT : [0-9]+;
ID : [a-zA-Z]+;
WS : [ \t\r\n]+ -> skip;
NL : '\r'? '\n';

LINE_COMMENT
    :   '//' ~[\r\n]* -> skip
    ;
