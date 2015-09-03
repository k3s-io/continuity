package commands

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	pb "github.com/stevvooe/continuity/proto"
)

var DumpCmd = &cobra.Command{
	Use:   "dump <manifest>",
	Short: "Dump the contents of the manifest in protobuf text format",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("please specify a manifest")
		}

		p, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatalf("error reading manifest: %v", err)
		}

		var bm pb.Manifest

		if err := proto.Unmarshal(p, &bm); err != nil {
			log.Fatalf("error unmarshaling manifest: %v", err)
		}

		// TODO(stevvooe): For now, just dump the text format. Turn this into
		// nice text output later.
		if err := proto.MarshalText(os.Stdout, &bm); err != nil {
			log.Fatalf("error dumping manifest: %v", err)
		}
	},
}