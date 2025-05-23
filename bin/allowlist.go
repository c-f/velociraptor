package main

// Default allow lists for server side VQL plugins

var (
	allowed_plugins = []string{
		"artifact_definitions",
		"batch",
		"chain",
		"client_delete",
		"clients",
		"clock",
		"column_filter",
		"combine",
		"delay",
		"delete_events",
		"delete_flow",
		"diff",
		"enumerate_flow",
		"fifo",
		"flatten",
		"flow_logs",
		"flow_results",
		"flows",
		"foreach",
		"glob",
		"hunt_delete",
		"hunt_flows",
		"hunt_results",
		"hunts",
		"if",
		"items",
		"monitoring",
		"notebook_delete",
		"olevba",
		"parallelize",
		"parse_csv",
		"parse_ese",
		"parse_ese_catalog",
		"parse_evtx",
		"parse_json_array",
		"parse_jsonl",
		"parse_lines",
		"parse_mft",
		"parse_ntfs_i30",
		"parse_ntfs_ranges",
		"parse_records_with_regex",
		"parse_recyclebin",
		"parse_usn",
		"plist",
		"prefetch",
		"query",
		"range",
		"read_file",
		"sample",
		"scope",
		"sequence",
		"source",
		"split_records",
		"stat",
		"switch",
		"timeline",
		"unzip",
		"uploads",
		"yara",
	}

	allowed_functions = []string{
		"add_client_monitoring",
		"all",
		"any",
		"array",
		"artifact_delete",
		"artifact_set",
		"atexit",
		"atoi",
		"base64decode",
		"base64encode",
		"basename",
		"cache",
		"cancel_flow",
		"cidr_contains",
		"client_create",
		"client_metadata",
		"client_set_metadata",
		"collect_client",
		"commandline_split",
		"compress",
		"count",
		"create_flow_download",
		"create_hunt_download",
		"crypto_rc4",
		"dict",
		"dirname",
		"encode",
		"entropy",
		"enumerate",
		"favorites_delete",
		"favorites_save",
		"file_store_delete",
		"filter",
		"format",
		"generate",
		"get",
		"get_client_monitoring",
		"get_flow",
		"grok",
		"gunzip",
		"hash",
		"humanize",
		"hunt",
		"hunt_add",
		"if",
		"import_collection",
		"int",
		"ip",
		"items",
		"join",
		"label",
		"len",
		"log",
		"lowcase",
		"lru",
		"lzxpress_decompress",
		"max",
		"memoize",
		"min",
		"now",
		"parse_binary",
		"parse_float",
		"parse_json",
		"parse_json_array",
		"parse_ntfs",
		"parse_pe",
		"parse_pkcs7",
		"parse_string_with_regex",
		"parse_x509",
		"parse_xml",
		"parse_yaml",
		"patch",
		"path_join",
		"path_split",
		"pathspec",
		"pipe",
		"plist",
		"rand",
		"rate",
		"read_file",
		"relpath",
		"repack",
		"remap",
		"rm_client_monitoring",
		"rot13",
		"scope",
		"serialize",
		"server_metadata",
		"server_set_metadata",
		"set",
		"set_client_monitoring",
		"sleep",
		"slice",
		"split",
		"str",
		"strip",
		"substr",
		"sum",
		"timeline_add",
		"timestamp",
		"to_dict",
		"unhex",
		"upcase",
		"upload",
		"url",
		"utf16",
		"utf16_encode",
		"uuid",
		"version",
		"whoami",
		"xor",
	}

	allowed_accessors = []string{
		"fs",
		"data",
		"scope",
		"collector",
		"bzip2",
		"gzip",
		"zip",
		"raw_reg",
		"mft",
	}
)
