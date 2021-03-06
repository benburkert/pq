package pq

const (
	t_bool                                  = 16
	t_bytea                                 = 17
	t_char                                  = 18
	t_name                                  = 19
	t_int8                                  = 20
	t_int2                                  = 21
	t_int2vector                            = 22
	t_int4                                  = 23
	t_regproc                               = 24
	t_text                                  = 25
	t_oid                                   = 26
	t_tid                                   = 27
	t_xid                                   = 28
	t_cid                                   = 29
	t_oidvector                             = 30
	t_pg_type                               = 71
	t_pg_attribute                          = 75
	t_pg_proc                               = 81
	t_pg_class                              = 83
	t_xml                                   = 142
	t__xml                                  = 143
	t_pg_node_tree                          = 194
	t_smgr                                  = 210
	t_point                                 = 600
	t_lseg                                  = 601
	t_path                                  = 602
	t_box                                   = 603
	t_polygon                               = 604
	t_line                                  = 628
	t__line                                 = 629
	t_float4                                = 700
	t_float8                                = 701
	t_abstime                               = 702
	t_reltime                               = 703
	t_tinterval                             = 704
	t_unknown                               = 705
	t_circle                                = 718
	t__circle                               = 719
	t_money                                 = 790
	t__money                                = 791
	t_macaddr                               = 829
	t_inet                                  = 869
	t_cidr                                  = 650
	t__bool                                 = 1000
	t__bytea                                = 1001
	t__char                                 = 1002
	t__name                                 = 1003
	t__int2                                 = 1005
	t__int2vector                           = 1006
	t__int4                                 = 1007
	t__regproc                              = 1008
	t__text                                 = 1009
	t__oid                                  = 1028
	t__tid                                  = 1010
	t__xid                                  = 1011
	t__cid                                  = 1012
	t__oidvector                            = 1013
	t__bpchar                               = 1014
	t__varchar                              = 1015
	t__int8                                 = 1016
	t__point                                = 1017
	t__lseg                                 = 1018
	t__path                                 = 1019
	t__box                                  = 1020
	t__float4                               = 1021
	t__float8                               = 1022
	t__abstime                              = 1023
	t__reltime                              = 1024
	t__tinterval                            = 1025
	t__polygon                              = 1027
	t_aclitem                               = 1033
	t__aclitem                              = 1034
	t__macaddr                              = 1040
	t__inet                                 = 1041
	t__cidr                                 = 651
	t__cstring                              = 1263
	t_bpchar                                = 1042
	t_varchar                               = 1043
	t_date                                  = 1082
	t_time                                  = 1083
	t_timestamp                             = 1114
	t__timestamp                            = 1115
	t__date                                 = 1182
	t__time                                 = 1183
	t_timestamptz                           = 1184
	t__timestamptz                          = 1185
	t_interval                              = 1186
	t__interval                             = 1187
	t__numeric                              = 1231
	t_timetz                                = 1266
	t__timetz                               = 1270
	t_bit                                   = 1560
	t__bit                                  = 1561
	t_varbit                                = 1562
	t__varbit                               = 1563
	t_numeric                               = 1700
	t_refcursor                             = 1790
	t__refcursor                            = 2201
	t_regprocedure                          = 2202
	t_regoper                               = 2203
	t_regoperator                           = 2204
	t_regclass                              = 2205
	t_regtype                               = 2206
	t__regprocedure                         = 2207
	t__regoper                              = 2208
	t__regoperator                          = 2209
	t__regclass                             = 2210
	t__regtype                              = 2211
	t_uuid                                  = 2950
	t__uuid                                 = 2951
	t_tsvector                              = 3614
	t_gtsvector                             = 3642
	t_tsquery                               = 3615
	t_regconfig                             = 3734
	t_regdictionary                         = 3769
	t__tsvector                             = 3643
	t__gtsvector                            = 3644
	t__tsquery                              = 3645
	t__regconfig                            = 3735
	t__regdictionary                        = 3770
	t_txid_snapshot                         = 2970
	t__txid_snapshot                        = 2949
	t_record                                = 2249
	t__record                               = 2287
	t_cstring                               = 2275
	t_any                                   = 2276
	t_anyarray                              = 2277
	t_void                                  = 2278
	t_trigger                               = 2279
	t_language_handler                      = 2280
	t_internal                              = 2281
	t_opaque                                = 2282
	t_anyelement                            = 2283
	t_anynonarray                           = 2776
	t_anyenum                               = 3500
	t_fdw_handler                           = 3115
	t_pg_attrdef                            = 10000
	t_pg_constraint                         = 10001
	t_pg_inherits                           = 10002
	t_pg_index                              = 10003
	t_pg_operator                           = 10004
	t_pg_opfamily                           = 10005
	t_pg_opclass                            = 10006
	t_pg_am                                 = 10117
	t_pg_amop                               = 10118
	t_pg_amproc                             = 10478
	t_pg_language                           = 10731
	t_pg_largeobject_metadata               = 10732
	t_pg_largeobject                        = 10733
	t_pg_aggregate                          = 10734
	t_pg_statistic                          = 10735
	t_pg_rewrite                            = 10736
	t_pg_trigger                            = 10737
	t_pg_description                        = 10738
	t_pg_cast                               = 10739
	t_pg_enum                               = 10936
	t_pg_namespace                          = 10937
	t_pg_conversion                         = 10938
	t_pg_depend                             = 10939
	t_pg_database                           = 1248
	t_pg_db_role_setting                    = 10940
	t_pg_tablespace                         = 10941
	t_pg_pltemplate                         = 10942
	t_pg_authid                             = 2842
	t_pg_auth_members                       = 2843
	t_pg_shdepend                           = 10943
	t_pg_shdescription                      = 10944
	t_pg_ts_config                          = 10945
	t_pg_ts_config_map                      = 10946
	t_pg_ts_dict                            = 10947
	t_pg_ts_parser                          = 10948
	t_pg_ts_template                        = 10949
	t_pg_extension                          = 10950
	t_pg_foreign_data_wrapper               = 10951
	t_pg_foreign_server                     = 10952
	t_pg_user_mapping                       = 10953
	t_pg_foreign_table                      = 10954
	t_pg_default_acl                        = 10955
	t_pg_seclabel                           = 10956
	t_pg_collation                          = 10957
	t_pg_toast_2604                         = 10958
	t_pg_toast_2606                         = 10959
	t_pg_toast_2609                         = 10960
	t_pg_toast_1255                         = 10961
	t_pg_toast_2618                         = 10962
	t_pg_toast_3596                         = 10963
	t_pg_toast_2619                         = 10964
	t_pg_toast_2620                         = 10965
	t_pg_toast_1262                         = 10966
	t_pg_toast_2396                         = 10967
	t_pg_toast_2964                         = 10968
	t_pg_roles                              = 10970
	t_pg_shadow                             = 10973
	t_pg_group                              = 10976
	t_pg_user                               = 10979
	t_pg_rules                              = 10982
	t_pg_views                              = 10986
	t_pg_tables                             = 10989
	t_pg_indexes                            = 10993
	t_pg_stats                              = 10997
	t_pg_locks                              = 11001
	t_pg_cursors                            = 11004
	t_pg_available_extensions               = 11007
	t_pg_available_extension_versions       = 11010
	t_pg_prepared_xacts                     = 11013
	t_pg_prepared_statements                = 11017
	t_pg_seclabels                          = 11020
	t_pg_settings                           = 11024
	t_pg_timezone_abbrevs                   = 11029
	t_pg_timezone_names                     = 11032
	t_pg_stat_all_tables                    = 11035
	t_pg_stat_xact_all_tables               = 11039
	t_pg_stat_sys_tables                    = 11043
	t_pg_stat_xact_sys_tables               = 11047
	t_pg_stat_user_tables                   = 11050
	t_pg_stat_xact_user_tables              = 11054
	t_pg_statio_all_tables                  = 11057
	t_pg_statio_sys_tables                  = 11061
	t_pg_statio_user_tables                 = 11064
	t_pg_stat_all_indexes                   = 11067
	t_pg_stat_sys_indexes                   = 11071
	t_pg_stat_user_indexes                  = 11074
	t_pg_statio_all_indexes                 = 11077
	t_pg_statio_sys_indexes                 = 11081
	t_pg_statio_user_indexes                = 11084
	t_pg_statio_all_sequences               = 11087
	t_pg_statio_sys_sequences               = 11090
	t_pg_statio_user_sequences              = 11093
	t_pg_stat_activity                      = 11096
	t_pg_stat_replication                   = 11099
	t_pg_stat_database                      = 11102
	t_pg_stat_database_conflicts            = 11105
	t_pg_stat_user_functions                = 11108
	t_pg_stat_xact_user_functions           = 11112
	t_pg_stat_bgwriter                      = 11116
	t_pg_user_mappings                      = 11119
	t_cardinal_number                       = 11669
	t_character_data                        = 11671
	t_sql_identifier                        = 11672
	t_information_schema_catalog_name       = 11674
	t_time_stamp                            = 11676
	t_yes_or_no                             = 11677
	t_applicable_roles                      = 11680
	t_administrable_role_authorizations     = 11684
	t_attributes                            = 11687
	t_character_sets                        = 11691
	t_check_constraint_routine_usage        = 11695
	t_check_constraints                     = 11699
	t_collations                            = 11703
	t_collation_character_set_applicability = 11706
	t_column_domain_usage                   = 11709
	t_column_privileges                     = 11713
	t_column_udt_usage                      = 11717
	t_columns                               = 11721
	t_constraint_column_usage               = 11725
	t_constraint_table_usage                = 11729
	t_domain_constraints                    = 11733
	t_domain_udt_usage                      = 11737
	t_domains                               = 11740
	t_enabled_roles                         = 11744
	t_key_column_usage                      = 11747
	t_parameters                            = 11751
	t_referential_constraints               = 11755
	t_role_column_grants                    = 11759
	t_routine_privileges                    = 11762
	t_role_routine_grants                   = 11766
	t_routines                              = 11769
	t_schemata                              = 11773
	t_sequences                             = 11776
	t_sql_features                          = 11780
	t_pg_toast_11779                        = 11782
	t_sql_implementation_info               = 11785
	t_pg_toast_11784                        = 11787
	t_sql_languages                         = 11790
	t_pg_toast_11789                        = 11792
	t_sql_packages                          = 11795
	t_pg_toast_11794                        = 11797
	t_sql_parts                             = 11800
	t_pg_toast_11799                        = 11802
	t_sql_sizing                            = 11805
	t_pg_toast_11804                        = 11807
	t_sql_sizing_profiles                   = 11810
	t_pg_toast_11809                        = 11812
	t_table_constraints                     = 11815
	t_table_privileges                      = 11819
	t_role_table_grants                     = 11823
	t_tables                                = 11826
	t_triggered_update_columns              = 11830
	t_triggers                              = 11834
	t_usage_privileges                      = 11838
	t_role_usage_grants                     = 11842
	t_view_column_usage                     = 11845
	t_view_routine_usage                    = 11849
	t_view_table_usage                      = 11853
	t_views                                 = 11857
	t_data_type_privileges                  = 11861
	t_element_types                         = 11865
	t__pg_foreign_data_wrappers             = 11869
	t_foreign_data_wrapper_options          = 11872
	t_foreign_data_wrappers                 = 11875
	t__pg_foreign_servers                   = 11878
	t_foreign_server_options                = 11882
	t_foreign_servers                       = 11885
	t__pg_foreign_tables                    = 11888
	t_foreign_table_options                 = 11892
	t_foreign_tables                        = 11895
	t__pg_user_mappings                     = 11898
	t_user_mapping_options                  = 11901
	t_user_mappings                         = 11905
	t_t                                     = 16806
	t__t                                    = 16805
	t_temp                                  = 16810
	t__temp                                 = 16809
)
