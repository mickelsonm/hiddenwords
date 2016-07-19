package hiddenwords

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCharLookup(t *testing.T) {
	Convey("handle when empty", t, func() {
		lookup, err := GetCharLookup("")
		So(lookup, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("handle valid character set", t, func() {
		lookup, err := GetCharLookup("abcdefghijklmnopqrstuvwxyz_")
		So(lookup, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestProcessText(t *testing.T) {
	Convey("handle no provided text", t, func() {
		lookup, _ := GetCharLookup("abcdefghijklmnopqrstuvwxyz_")
		processed, err := ProcessText("", lookup)
		So(processed, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("handle valid text and lookup", t, func() {
		lookup, _ := GetCharLookup("abcdefghijklmnopqrstuvwxyz_")
		searchText := `epqiiqwdiwgyka_vsqtsujeqqicnhyivo_sigwasmkwgsih_akl_gtnkhgikgveidpmt
    qybpxpnnpbxkwpisgjmdzgh_ojysbtsnsvxvuhguocp_qc_vouxqmg_cetlpmounxnvg
    ldcpem_jodnmklgonocekdkjwkdoilajk_nxujykigsolengqmnqofpseqaamvpsooga
    spyhoojennefwvljpvsqtgnceg_hsowqvycjkuxdtfbxfloewkphmvkftjlsasvwid_u
    qcsgn_ypiqjytygiwyziqdjpxgpuunymadnclpdlmmulitsnqlwciotbmyfuummjynne
    slnit_lpykdafkpydzkntbud_gigjgmu_uqjjmdzpwteodjpuzndxaqmsjdjjamnwoes
    ajcffkaaoilpyydlkyxauagfcjbabapax_ndlgtpwnud_jpnkiokviqjhyopmjtgtbyo
    iyfbjdhknimlah_cxfzwspqoscffiyvabtjjuc_liaqbcuomuytdqfy_xaixiiqqdpds
    uuimzh_ywwcmodxhfxjplyixotjkeawauxltekptuieekpbokbanumffatbtiacnywhw
    iqxebnosninpzfjmatvnyuspyeu_ziapvogconld_cxfcytkcp_bvsppz_dw_ndlpkhf
    zdlxbo_vaflmailjvccgsuclyhojganjqxzmqflpze_hqhlul_ybaagtiuokbzaxhmec
    olsptiexvvmhbdoelgmcffulcebhlyzd_m_qxkbfvnxykdudpxefsm_aqpqtnhxvswht
    owqnbm_mgejjpyumm_mqbkiuulanbmzllmuqlfftmcxtybmijfuwaknefhekwgujpjqg
    leu_sjtbszotcygiclkwcbmnvgsoqaqqkkgeaslhvfbtlgpnxgpzxp_vyjinlwwfbvtn
    twogmnpxghabpxxgzlyirrrrrbbcrrrnbjpcrrrqykhrrrscarrrdnlxrrrrtudrrrr_
    ntrbyrqlddbycypcccqongpgexhnabavrmebeofrxsnrilprveetxaranjyfmrisrewp
    r_y_lgsrsedbn_rfrieusemhpfa_plkifjipvwaqvnenrrrzybsrbeurbhfrvrrzghr_
    zpgiyrrrqsnnrrrbhvdrrrqkpdrraqvkeueszfpkj_fm_claw_oetbgurbdocb_rsnzr
    cyvrvnrvaurbscimurtbriikrfdjlizribdjwkror_gnlzmshwccqcx_huaafbvituxo
    ru_hohxwrrrhnbttrrriyyirrrnibricrxftrrrrvqvrrrrhjorehroldibsmquelwvy
    jebkolbbnauompgqdhlbnsfbbdiudoeibwstdg_acsazhtgfufidogmyvtya_dfwihto
    elucbtlcbaijlcuhfvhesgluiwttsdnqqshnoqumccyqtko_zh_fii_wlsspysdqdpad
    fvfewlsojavmuaixyxpw_xcwxuatceosdqgmsbbagjmmblouvnywmqqakmmtuasfovol
    _ogksdukwp_fkxuh_vfhuhfyfvvfqhqxecxsoctcqgpianhtnkbqlltwyhxotfksoewm
    elxobjgwlyfaeoxsfohhguidoftbsainwovvglynsgjixon_nvuwflsfbca_xnnesvco
    mceh_gigjxpllckcooagidcpbqxtnejlnlsccocuvcvge_fvjjbyqdkjceia_mkcvbzl
    zwlxbdjihvpmdcvmssuvktwiqbeivtieol_bu_huumzmlxx_kd_vksmohgzl_fxwfdue
    lqgfkgzxciwmuduozfbaxstxkwegescggkpxfpeenhb_whqhethcateqdvnxhpt__bja
    _uiyxchmfkblmdwtyp_ktontmufw_isdflelsbgjizxvqbciuadfxxjaqbluofkgkkkh
    jbvohisfla_cspbmuezqohnyijyimwgdeszutgnaoagbhku_wwdtylbbiyvbpoumgyid
    w_xwg_fkogabccip_wouclnjcgdpwwxxvvvwkmmbgfeactbcksxqovqthtjfjghijwwh
    ydfieyssbjtfqgqyjnmwfpesljmwapvbptucadontbobnspch_i_dxheklulncdsdnic
    bnjjjedkaokw_ahcolvbcnmqtoakonpgzjufqlnn_uve_uumaufjasfvfcv_cbcuk_hd
    zigkahchzfqjphjwcbjwmozyodhu_tsqtafwidgmc_snhhkleyvmzdtawdodzfmekuee
    mnshz_xz`

		processed, err := ProcessText(searchText, lookup)
		So(processed, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(processed[0].Text, ShouldEqual, "b")
		So(processed[0].Count, ShouldEqual, 100)
		So(processed[9].Text, ShouldEqual, "s")
		So(processed[9].Count, ShouldEqual, 91)
	})
}

func TestGetAnswer(t *testing.T) {
	Convey("should have a valid answer", t, func() {
		lookup, _ := GetCharLookup("abcdefghijklmnopqrstuvwxyz_")
		searchText := `epqiiqwdiwgyka_vsqtsujeqqicnhyivo_sigwasmkwgsih_akl_gtnkhgikgveidpmt
    qybpxpnnpbxkwpisgjmdzgh_ojysbtsnsvxvuhguocp_qc_vouxqmg_cetlpmounxnvg
    ldcpem_jodnmklgonocekdkjwkdoilajk_nxujykigsolengqmnqofpseqaamvpsooga
    spyhoojennefwvljpvsqtgnceg_hsowqvycjkuxdtfbxfloewkphmvkftjlsasvwid_u
    qcsgn_ypiqjytygiwyziqdjpxgpuunymadnclpdlmmulitsnqlwciotbmyfuummjynne
    slnit_lpykdafkpydzkntbud_gigjgmu_uqjjmdzpwteodjpuzndxaqmsjdjjamnwoes
    ajcffkaaoilpyydlkyxauagfcjbabapax_ndlgtpwnud_jpnkiokviqjhyopmjtgtbyo
    iyfbjdhknimlah_cxfzwspqoscffiyvabtjjuc_liaqbcuomuytdqfy_xaixiiqqdpds
    uuimzh_ywwcmodxhfxjplyixotjkeawauxltekptuieekpbokbanumffatbtiacnywhw
    iqxebnosninpzfjmatvnyuspyeu_ziapvogconld_cxfcytkcp_bvsppz_dw_ndlpkhf
    zdlxbo_vaflmailjvccgsuclyhojganjqxzmqflpze_hqhlul_ybaagtiuokbzaxhmec
    olsptiexvvmhbdoelgmcffulcebhlyzd_m_qxkbfvnxykdudpxefsm_aqpqtnhxvswht
    owqnbm_mgejjpyumm_mqbkiuulanbmzllmuqlfftmcxtybmijfuwaknefhekwgujpjqg
    leu_sjtbszotcygiclkwcbmnvgsoqaqqkkgeaslhvfbtlgpnxgpzxp_vyjinlwwfbvtn
    twogmnpxghabpxxgzlyirrrrrbbcrrrnbjpcrrrqykhrrrscarrrdnlxrrrrtudrrrr_
    ntrbyrqlddbycypcccqongpgexhnabavrmebeofrxsnrilprveetxaranjyfmrisrewp
    r_y_lgsrsedbn_rfrieusemhpfa_plkifjipvwaqvnenrrrzybsrbeurbhfrvrrzghr_
    zpgiyrrrqsnnrrrbhvdrrrqkpdrraqvkeueszfpkj_fm_claw_oetbgurbdocb_rsnzr
    cyvrvnrvaurbscimurtbriikrfdjlizribdjwkror_gnlzmshwccqcx_huaafbvituxo
    ru_hohxwrrrhnbttrrriyyirrrnibricrxftrrrrvqvrrrrhjorehroldibsmquelwvy
    jebkolbbnauompgqdhlbnsfbbdiudoeibwstdg_acsazhtgfufidogmyvtya_dfwihto
    elucbtlcbaijlcuhfvhesgluiwttsdnqqshnoqumccyqtko_zh_fii_wlsspysdqdpad
    fvfewlsojavmuaixyxpw_xcwxuatceosdqgmsbbagjmmblouvnywmqqakmmtuasfovol
    _ogksdukwp_fkxuh_vfhuhfyfvvfqhqxecxsoctcqgpianhtnkbqlltwyhxotfksoewm
    elxobjgwlyfaeoxsfohhguidoftbsainwovvglynsgjixon_nvuwflsfbca_xnnesvco
    mceh_gigjxpllckcooagidcpbqxtnejlnlsccocuvcvge_fvjjbyqdkjceia_mkcvbzl
    zwlxbdjihvpmdcvmssuvktwiqbeivtieol_bu_huumzmlxx_kd_vksmohgzl_fxwfdue
    lqgfkgzxciwmuduozfbaxstxkwegescggkpxfpeenhb_whqhethcateqdvnxhpt__bja
    _uiyxchmfkblmdwtyp_ktontmufw_isdflelsbgjizxvqbciuadfxxjaqbluofkgkkkh
    jbvohisfla_cspbmuezqohnyijyimwgdeszutgnaoagbhku_wwdtylbbiyvbpoumgyid
    w_xwg_fkogabccip_wouclnjcgdpwwxxvvvwkmmbgfeactbcksxqovqthtjfjghijwwh
    ydfieyssbjtfqgqyjnmwfpesljmwapvbptucadontbobnspch_i_dxheklulncdsdnic
    bnjjjedkaokw_ahcolvbcnmqtoakonpgzjufqlnn_uve_uumaufjasfvfcv_cbcuk_hd
    zigkahchzfqjphjwcbjwmozyodhu_tsqtafwidgmc_snhhkleyvmzdtawdodzfmekuee
    mnshz_xz`

		processed, err := ProcessText(searchText, lookup)
		So(processed, ShouldNotBeNil)
		So(err, ShouldBeNil)

		answer := processed.GetAnswer()
		So(answer, ShouldEqual, "binoculars")
	})
}
