grammar Nugget2;

prog: (define_assign | NL
        )*
        EOF
;


define_assign:   define |
        assign
;

define: ID nugget_type LISTOP?;

assign: ID '=' STRING asType ('|' nuggetaction)* |
        ID '=' ID ('|' nuggetaction)*
;

asType: 'as' nugget_type;

nugget_type: 'string'  |
      'sha1'     |
      'md5'      |
      'ntfs'     |
      'file'     |
      'packet'   |
      'exifinfo' |
;

nuggetaction: 'extract' ID |
        'sha1' |
        'md5'  |
;

LISTOP: '[]';

INT : [0-9]+;
ID : [a-zA-Z]+;
STRING: '"' ('""'|~'"')* '"';

WS : [ \t\r\n]+ -> skip;
NL : '\r'? '\n';

LINE_COMMENT
    :   '//' ~[\r\n]* -> skip
    ;
