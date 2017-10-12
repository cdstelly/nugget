grammar Nugget;

@header {
    // import "../NTypes"
}

prog: (         define_assign |
       operation_on_singleton |
         singleton_var )*
       EOF
;

define_assign:   define |
                 define_tuple |
                 assign
;

define: ID nugget_type LISTOP? ;

define_tuple: ID 'tuple[' (','? nugget_type)+ ']' LISTOP?;


assign: ID '=' STRING ('|' nugget_action)* |
        ID '=' ID ('|' nugget_action)*
;

operation_on_singleton: singleton_op '(' ID ')';

singleton_op: ('type' | 'print' | 'size' | 'typex' | 'printx' | 'raw');

singleton_var: ID;

nugget_type:
      'string'     |
      'sha1'       |
      'md5'        |
      'ntfs'       |
      'file'       |
      'packet'     |
      'pcap'       |
      'exifinfo'   |
      'datetime'   |
      'memory'     |
      'listof-md5' |
      'listof-sha1'|
      'listof-sha256'
      ;


nugget_action: action_word ;

action_word:
    filter            |
    'extract' asType  |
    'sort'    byField |
    'sha1'            |
    'md5'             |
	'sha256'          |
	'getGetRequests'  |
	'diskinfo'        |
	'union'    ID     |
	'pslist'          |
	'grep'     STRING |
	'%%%'
;

asType: 'as' nugget_type (byteOffsetSize)?;
byField:'by' ID;

byteOffsetSize : '['INT ',' INT ']';

filter : 'filter' filter_term (',' filter_term)*;
filter_term: ID COMPOP STRING;

COMPOP: ('>' | '<' | '>=' | '<=' | '==');
LISTOP: '[]';

INT : [0-9]+;
ID : [a-zA-Z]+;
STRING: '"' ('""'|~'"')* '"';

WS : [ \t\r\n]+ -> skip;
NL : '\r'? '\n';

LINE_COMMENT: '//' ~[\r\n]* -> skip;
