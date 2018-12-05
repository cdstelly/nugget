@0xc42f21aa6a1470d6

struct NetFileinfo { 
	id @0 :Text; 
	filenames @0 :List(Text);
	createtime @0 :Int64;
	modifytime @0 :Int64; 
	accesstime @0 :Int64;
	emodifytime @0 :Int64;
	fflags @0 :Text;
	flags @0 :Text;
	filesize @0 :UInt64;
	reconstructeddata @0 :Bytes;
	beenreconstructed @0 :Bool;
}


interface NetFileinfoService {
	getFileinfo @0 (query :Text) -> (result :Text);
}
