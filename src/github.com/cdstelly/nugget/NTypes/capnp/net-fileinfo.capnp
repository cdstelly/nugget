using Go = import "/go.capnp";
$Go.package("NTypes");
$Go.import("github.com/cdstelly/nugget/NTypes");

@0xc42f21aa6a1470d6;

struct NetFileinfo { 
	id @0 :Text; 
	filenames @1 :List(Text);
	createtime @2 :Int64;
	modifytime @3 :Int64; 
	accesstime @4 :Int64;
	emodifytime @5 :Int64;
	fflags @6 :Text;
	flags @7 :Text;
	filesize @8 :UInt64;
	reconstructeddata @9 :Data;
	beenreconstructed @10 :Bool;
}


interface NetFileinfoService {
	getFileinfo @0 (query :Text) -> (result :NetFileinfo);
}
