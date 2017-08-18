grammar Nugget2;

prog: (define_assign | NL
        )*
        EOF
;


define_assign:   define |
                 assign |
                 singleton_var
;

define: ID nugget_type LISTOP?;

assign: ID '=' STRING asType ('|' nugget_action)* |
        ID '=' ID ('|' nugget_action)*
;

singleton_var: ID;

asType: 'as' nugget_type;

nugget_type: 'string'  |
      'sha1'     |
      'md5'      |
      'ntfs'     |
      'file'     |
      'packet'   |
      'exifinfo'
;

nugget_action: action_word (ID)?
;

action_word: 'extract' |
             'sha1'    |
             'md5'
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
