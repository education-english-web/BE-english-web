package fileutil

import "testing"

func TestExtractFromFilename(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name     string
		args     args
		wantName string
		wantExt  string
	}{
		{
			name: "Case 1",
			args: args{
				filename: "test/test/filename.ext",
			},
			wantName: "filename",
			wantExt:  "ext",
		},
		{
			name: "Case 1a",
			args: args{
				filename: "test/test/filename.EXT",
			},
			wantName: "filename",
			wantExt:  "ext",
		},
		{
			name: "Case 2",
			args: args{
				filename: "test/test/filename",
			},
			wantName: "filename",
			wantExt:  "",
		},
		{
			name: "Case 3",
			args: args{
				filename: "test/test/.ext",
			},
			wantName: "",
			wantExt:  "ext",
		},
		{
			name: "Case 4",
			args: args{
				filename: "test/test/a..ext",
			},
			wantName: "a.",
			wantExt:  "ext",
		},
		{
			name: "Case 4",
			args: args{
				filename: "test/test/filename.",
			},
			wantName: "filename",
			wantExt:  "",
		},
		{
			name: "Case 5",
			args: args{
				filename: "test/test/.",
			},
			wantName: "",
			wantExt:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotExt := ExtractFromFilename(tt.args.filename)
			if gotName != tt.wantName {
				t.Errorf("ExtractFromFilename() gotName = %v, want %v", gotName, tt.wantName)
			}

			if gotExt != tt.wantExt {
				t.Errorf("ExtractFromFilename() gotExt = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}
