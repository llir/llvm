// Code generated by "stringer -linecomment -type DwarfOp"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DwarfOpAddr-3]
	_ = x[DwarfOpDeref-6]
	_ = x[DwarfOpConst1u-8]
	_ = x[DwarfOpConst1s-9]
	_ = x[DwarfOpConst2u-10]
	_ = x[DwarfOpConst2s-11]
	_ = x[DwarfOpConst4u-12]
	_ = x[DwarfOpConst4s-13]
	_ = x[DwarfOpConst8u-14]
	_ = x[DwarfOpConst8s-15]
	_ = x[DwarfOpConstu-16]
	_ = x[DwarfOpConsts-17]
	_ = x[DwarfOpDup-18]
	_ = x[DwarfOpDrop-19]
	_ = x[DwarfOpOver-20]
	_ = x[DwarfOpPick-21]
	_ = x[DwarfOpSwap-22]
	_ = x[DwarfOpRot-23]
	_ = x[DwarfOpXderef-24]
	_ = x[DwarfOpAbs-25]
	_ = x[DwarfOpAnd-26]
	_ = x[DwarfOpDiv-27]
	_ = x[DwarfOpMinus-28]
	_ = x[DwarfOpMod-29]
	_ = x[DwarfOpMul-30]
	_ = x[DwarfOpNeg-31]
	_ = x[DwarfOpNot-32]
	_ = x[DwarfOpOr-33]
	_ = x[DwarfOpPlus-34]
	_ = x[DwarfOpPlusUconst-35]
	_ = x[DwarfOpShl-36]
	_ = x[DwarfOpShr-37]
	_ = x[DwarfOpShra-38]
	_ = x[DwarfOpXor-39]
	_ = x[DwarfOpBra-40]
	_ = x[DwarfOpEq-41]
	_ = x[DwarfOpGe-42]
	_ = x[DwarfOpGt-43]
	_ = x[DwarfOpLe-44]
	_ = x[DwarfOpLt-45]
	_ = x[DwarfOpNe-46]
	_ = x[DwarfOpSkip-47]
	_ = x[DwarfOpLit0-48]
	_ = x[DwarfOpLit1-49]
	_ = x[DwarfOpLit2-50]
	_ = x[DwarfOpLit3-51]
	_ = x[DwarfOpLit4-52]
	_ = x[DwarfOpLit5-53]
	_ = x[DwarfOpLit6-54]
	_ = x[DwarfOpLit7-55]
	_ = x[DwarfOpLit8-56]
	_ = x[DwarfOpLit9-57]
	_ = x[DwarfOpLit10-58]
	_ = x[DwarfOpLit11-59]
	_ = x[DwarfOpLit12-60]
	_ = x[DwarfOpLit13-61]
	_ = x[DwarfOpLit14-62]
	_ = x[DwarfOpLit15-63]
	_ = x[DwarfOpLit16-64]
	_ = x[DwarfOpLit17-65]
	_ = x[DwarfOpLit18-66]
	_ = x[DwarfOpLit19-67]
	_ = x[DwarfOpLit20-68]
	_ = x[DwarfOpLit21-69]
	_ = x[DwarfOpLit22-70]
	_ = x[DwarfOpLit23-71]
	_ = x[DwarfOpLit24-72]
	_ = x[DwarfOpLit25-73]
	_ = x[DwarfOpLit26-74]
	_ = x[DwarfOpLit27-75]
	_ = x[DwarfOpLit28-76]
	_ = x[DwarfOpLit29-77]
	_ = x[DwarfOpLit30-78]
	_ = x[DwarfOpLit31-79]
	_ = x[DwarfOpReg0-80]
	_ = x[DwarfOpReg1-81]
	_ = x[DwarfOpReg2-82]
	_ = x[DwarfOpReg3-83]
	_ = x[DwarfOpReg4-84]
	_ = x[DwarfOpReg5-85]
	_ = x[DwarfOpReg6-86]
	_ = x[DwarfOpReg7-87]
	_ = x[DwarfOpReg8-88]
	_ = x[DwarfOpReg9-89]
	_ = x[DwarfOpReg10-90]
	_ = x[DwarfOpReg11-91]
	_ = x[DwarfOpReg12-92]
	_ = x[DwarfOpReg13-93]
	_ = x[DwarfOpReg14-94]
	_ = x[DwarfOpReg15-95]
	_ = x[DwarfOpReg16-96]
	_ = x[DwarfOpReg17-97]
	_ = x[DwarfOpReg18-98]
	_ = x[DwarfOpReg19-99]
	_ = x[DwarfOpReg20-100]
	_ = x[DwarfOpReg21-101]
	_ = x[DwarfOpReg22-102]
	_ = x[DwarfOpReg23-103]
	_ = x[DwarfOpReg24-104]
	_ = x[DwarfOpReg25-105]
	_ = x[DwarfOpReg26-106]
	_ = x[DwarfOpReg27-107]
	_ = x[DwarfOpReg28-108]
	_ = x[DwarfOpReg29-109]
	_ = x[DwarfOpReg30-110]
	_ = x[DwarfOpReg31-111]
	_ = x[DwarfOpBreg0-112]
	_ = x[DwarfOpBreg1-113]
	_ = x[DwarfOpBreg2-114]
	_ = x[DwarfOpBreg3-115]
	_ = x[DwarfOpBreg4-116]
	_ = x[DwarfOpBreg5-117]
	_ = x[DwarfOpBreg6-118]
	_ = x[DwarfOpBreg7-119]
	_ = x[DwarfOpBreg8-120]
	_ = x[DwarfOpBreg9-121]
	_ = x[DwarfOpBreg10-122]
	_ = x[DwarfOpBreg11-123]
	_ = x[DwarfOpBreg12-124]
	_ = x[DwarfOpBreg13-125]
	_ = x[DwarfOpBreg14-126]
	_ = x[DwarfOpBreg15-127]
	_ = x[DwarfOpBreg16-128]
	_ = x[DwarfOpBreg17-129]
	_ = x[DwarfOpBreg18-130]
	_ = x[DwarfOpBreg19-131]
	_ = x[DwarfOpBreg20-132]
	_ = x[DwarfOpBreg21-133]
	_ = x[DwarfOpBreg22-134]
	_ = x[DwarfOpBreg23-135]
	_ = x[DwarfOpBreg24-136]
	_ = x[DwarfOpBreg25-137]
	_ = x[DwarfOpBreg26-138]
	_ = x[DwarfOpBreg27-139]
	_ = x[DwarfOpBreg28-140]
	_ = x[DwarfOpBreg29-141]
	_ = x[DwarfOpBreg30-142]
	_ = x[DwarfOpBreg31-143]
	_ = x[DwarfOpRegx-144]
	_ = x[DwarfOpFbreg-145]
	_ = x[DwarfOpBregx-146]
	_ = x[DwarfOpPiece-147]
	_ = x[DwarfOpDerefSize-148]
	_ = x[DwarfOpXderefSize-149]
	_ = x[DwarfOpNop-150]
	_ = x[DwarfOpPushObjectAddress-151]
	_ = x[DwarfOpCall2-152]
	_ = x[DwarfOpCall4-153]
	_ = x[DwarfOpCallRef-154]
	_ = x[DwarfOpFormTLSAddress-155]
	_ = x[DwarfOpCallFrameCFA-156]
	_ = x[DwarfOpBitPiece-157]
	_ = x[DwarfOpImplicitValue-158]
	_ = x[DwarfOpStackValue-159]
	_ = x[DwarfOpImplicitPointer-160]
	_ = x[DwarfOpAddrx-161]
	_ = x[DwarfOpConstx-162]
	_ = x[DwarfOpEntryValue-163]
	_ = x[DwarfOpConstType-164]
	_ = x[DwarfOpRegvalType-165]
	_ = x[DwarfOpDerefType-166]
	_ = x[DwarfOpXderefType-167]
	_ = x[DwarfOpConvert-168]
	_ = x[DwarfOpReinterpret-169]
	_ = x[DwarfOpGNUPushTLSAddress-224]
	_ = x[DwarfOpHPIsValue-225]
	_ = x[DwarfOpHPFltConst4-226]
	_ = x[DwarfOpHPFltConst8-227]
	_ = x[DwarfOpHPModRange-228]
	_ = x[DwarfOpHPUnmodRange-229]
	_ = x[DwarfOpHPTLS-230]
	_ = x[DwarfOpIntelBitPiece-232]
	_ = x[DwarfOpWASMLocation-237]
	_ = x[DwarfOpWASMLocationInt-238]
	_ = x[DwarfOpAppleUninit-240]
	_ = x[DwarfOpGNUEntryValue-243]
	_ = x[DwarfOpPGIOmpThreadNum-248]
	_ = x[DwarfOpGNUAddrIndex-251]
	_ = x[DwarfOpGNUConstIndex-252]
	_ = x[DwarfOpLLVMFragment-4096]
	_ = x[DwarfOpLLVMConvert-4097]
	_ = x[DwarfOpLLVMTagOffset-4098]
	_ = x[DwarfOpLLVMEntryValue-4099]
	_ = x[DwarfOpLLVMImplicitPointer-4100]
	_ = x[DwarfOpLLVMArg-4101]
}

const _DwarfOp_name = "DW_OP_addrDW_OP_derefDW_OP_const1uDW_OP_const1sDW_OP_const2uDW_OP_const2sDW_OP_const4uDW_OP_const4sDW_OP_const8uDW_OP_const8sDW_OP_constuDW_OP_constsDW_OP_dupDW_OP_dropDW_OP_overDW_OP_pickDW_OP_swapDW_OP_rotDW_OP_xderefDW_OP_absDW_OP_andDW_OP_divDW_OP_minusDW_OP_modDW_OP_mulDW_OP_negDW_OP_notDW_OP_orDW_OP_plusDW_OP_plus_uconstDW_OP_shlDW_OP_shrDW_OP_shraDW_OP_xorDW_OP_braDW_OP_eqDW_OP_geDW_OP_gtDW_OP_leDW_OP_ltDW_OP_neDW_OP_skipDW_OP_lit0DW_OP_lit1DW_OP_lit2DW_OP_lit3DW_OP_lit4DW_OP_lit5DW_OP_lit6DW_OP_lit7DW_OP_lit8DW_OP_lit9DW_OP_lit10DW_OP_lit11DW_OP_lit12DW_OP_lit13DW_OP_lit14DW_OP_lit15DW_OP_lit16DW_OP_lit17DW_OP_lit18DW_OP_lit19DW_OP_lit20DW_OP_lit21DW_OP_lit22DW_OP_lit23DW_OP_lit24DW_OP_lit25DW_OP_lit26DW_OP_lit27DW_OP_lit28DW_OP_lit29DW_OP_lit30DW_OP_lit31DW_OP_reg0DW_OP_reg1DW_OP_reg2DW_OP_reg3DW_OP_reg4DW_OP_reg5DW_OP_reg6DW_OP_reg7DW_OP_reg8DW_OP_reg9DW_OP_reg10DW_OP_reg11DW_OP_reg12DW_OP_reg13DW_OP_reg14DW_OP_reg15DW_OP_reg16DW_OP_reg17DW_OP_reg18DW_OP_reg19DW_OP_reg20DW_OP_reg21DW_OP_reg22DW_OP_reg23DW_OP_reg24DW_OP_reg25DW_OP_reg26DW_OP_reg27DW_OP_reg28DW_OP_reg29DW_OP_reg30DW_OP_reg31DW_OP_breg0DW_OP_breg1DW_OP_breg2DW_OP_breg3DW_OP_breg4DW_OP_breg5DW_OP_breg6DW_OP_breg7DW_OP_breg8DW_OP_breg9DW_OP_breg10DW_OP_breg11DW_OP_breg12DW_OP_breg13DW_OP_breg14DW_OP_breg15DW_OP_breg16DW_OP_breg17DW_OP_breg18DW_OP_breg19DW_OP_breg20DW_OP_breg21DW_OP_breg22DW_OP_breg23DW_OP_breg24DW_OP_breg25DW_OP_breg26DW_OP_breg27DW_OP_breg28DW_OP_breg29DW_OP_breg30DW_OP_breg31DW_OP_regxDW_OP_fbregDW_OP_bregxDW_OP_pieceDW_OP_deref_sizeDW_OP_xderef_sizeDW_OP_nopDW_OP_push_object_addressDW_OP_call2DW_OP_call4DW_OP_call_refDW_OP_form_tls_addressDW_OP_call_frame_cfaDW_OP_bit_pieceDW_OP_implicit_valueDW_OP_stack_valueDW_OP_implicit_pointerDW_OP_addrxDW_OP_constxDW_OP_entry_valueDW_OP_const_typeDW_OP_regval_typeDW_OP_deref_typeDW_OP_xderef_typeDW_OP_convertDW_OP_reinterpretDW_OP_GNU_push_tls_addressDW_OP_HP_is_valueDW_OP_HP_fltconst4DW_OP_HP_fltconst8DW_OP_HP_mod_rangeDW_OP_HP_unmod_rangeDW_OP_HP_tlsDW_OP_INTEL_bit_pieceDW_OP_WASM_locationDW_OP_WASM_location_intDW_OP_APPLE_uninitDW_OP_GNU_entry_valueDW_OP_PGI_omp_thread_numDW_OP_GNU_addr_indexDW_OP_GNU_const_indexDW_OP_LLVM_fragmentDW_OP_LLVM_convertDW_OP_LLVM_tag_offsetDW_OP_LLVM_entry_valueDW_OP_LLVM_implicit_pointerDW_OP_LLVM_arg"

var _DwarfOp_map = map[DwarfOp]string{
	3:    _DwarfOp_name[0:10],
	6:    _DwarfOp_name[10:21],
	8:    _DwarfOp_name[21:34],
	9:    _DwarfOp_name[34:47],
	10:   _DwarfOp_name[47:60],
	11:   _DwarfOp_name[60:73],
	12:   _DwarfOp_name[73:86],
	13:   _DwarfOp_name[86:99],
	14:   _DwarfOp_name[99:112],
	15:   _DwarfOp_name[112:125],
	16:   _DwarfOp_name[125:137],
	17:   _DwarfOp_name[137:149],
	18:   _DwarfOp_name[149:158],
	19:   _DwarfOp_name[158:168],
	20:   _DwarfOp_name[168:178],
	21:   _DwarfOp_name[178:188],
	22:   _DwarfOp_name[188:198],
	23:   _DwarfOp_name[198:207],
	24:   _DwarfOp_name[207:219],
	25:   _DwarfOp_name[219:228],
	26:   _DwarfOp_name[228:237],
	27:   _DwarfOp_name[237:246],
	28:   _DwarfOp_name[246:257],
	29:   _DwarfOp_name[257:266],
	30:   _DwarfOp_name[266:275],
	31:   _DwarfOp_name[275:284],
	32:   _DwarfOp_name[284:293],
	33:   _DwarfOp_name[293:301],
	34:   _DwarfOp_name[301:311],
	35:   _DwarfOp_name[311:328],
	36:   _DwarfOp_name[328:337],
	37:   _DwarfOp_name[337:346],
	38:   _DwarfOp_name[346:356],
	39:   _DwarfOp_name[356:365],
	40:   _DwarfOp_name[365:374],
	41:   _DwarfOp_name[374:382],
	42:   _DwarfOp_name[382:390],
	43:   _DwarfOp_name[390:398],
	44:   _DwarfOp_name[398:406],
	45:   _DwarfOp_name[406:414],
	46:   _DwarfOp_name[414:422],
	47:   _DwarfOp_name[422:432],
	48:   _DwarfOp_name[432:442],
	49:   _DwarfOp_name[442:452],
	50:   _DwarfOp_name[452:462],
	51:   _DwarfOp_name[462:472],
	52:   _DwarfOp_name[472:482],
	53:   _DwarfOp_name[482:492],
	54:   _DwarfOp_name[492:502],
	55:   _DwarfOp_name[502:512],
	56:   _DwarfOp_name[512:522],
	57:   _DwarfOp_name[522:532],
	58:   _DwarfOp_name[532:543],
	59:   _DwarfOp_name[543:554],
	60:   _DwarfOp_name[554:565],
	61:   _DwarfOp_name[565:576],
	62:   _DwarfOp_name[576:587],
	63:   _DwarfOp_name[587:598],
	64:   _DwarfOp_name[598:609],
	65:   _DwarfOp_name[609:620],
	66:   _DwarfOp_name[620:631],
	67:   _DwarfOp_name[631:642],
	68:   _DwarfOp_name[642:653],
	69:   _DwarfOp_name[653:664],
	70:   _DwarfOp_name[664:675],
	71:   _DwarfOp_name[675:686],
	72:   _DwarfOp_name[686:697],
	73:   _DwarfOp_name[697:708],
	74:   _DwarfOp_name[708:719],
	75:   _DwarfOp_name[719:730],
	76:   _DwarfOp_name[730:741],
	77:   _DwarfOp_name[741:752],
	78:   _DwarfOp_name[752:763],
	79:   _DwarfOp_name[763:774],
	80:   _DwarfOp_name[774:784],
	81:   _DwarfOp_name[784:794],
	82:   _DwarfOp_name[794:804],
	83:   _DwarfOp_name[804:814],
	84:   _DwarfOp_name[814:824],
	85:   _DwarfOp_name[824:834],
	86:   _DwarfOp_name[834:844],
	87:   _DwarfOp_name[844:854],
	88:   _DwarfOp_name[854:864],
	89:   _DwarfOp_name[864:874],
	90:   _DwarfOp_name[874:885],
	91:   _DwarfOp_name[885:896],
	92:   _DwarfOp_name[896:907],
	93:   _DwarfOp_name[907:918],
	94:   _DwarfOp_name[918:929],
	95:   _DwarfOp_name[929:940],
	96:   _DwarfOp_name[940:951],
	97:   _DwarfOp_name[951:962],
	98:   _DwarfOp_name[962:973],
	99:   _DwarfOp_name[973:984],
	100:  _DwarfOp_name[984:995],
	101:  _DwarfOp_name[995:1006],
	102:  _DwarfOp_name[1006:1017],
	103:  _DwarfOp_name[1017:1028],
	104:  _DwarfOp_name[1028:1039],
	105:  _DwarfOp_name[1039:1050],
	106:  _DwarfOp_name[1050:1061],
	107:  _DwarfOp_name[1061:1072],
	108:  _DwarfOp_name[1072:1083],
	109:  _DwarfOp_name[1083:1094],
	110:  _DwarfOp_name[1094:1105],
	111:  _DwarfOp_name[1105:1116],
	112:  _DwarfOp_name[1116:1127],
	113:  _DwarfOp_name[1127:1138],
	114:  _DwarfOp_name[1138:1149],
	115:  _DwarfOp_name[1149:1160],
	116:  _DwarfOp_name[1160:1171],
	117:  _DwarfOp_name[1171:1182],
	118:  _DwarfOp_name[1182:1193],
	119:  _DwarfOp_name[1193:1204],
	120:  _DwarfOp_name[1204:1215],
	121:  _DwarfOp_name[1215:1226],
	122:  _DwarfOp_name[1226:1238],
	123:  _DwarfOp_name[1238:1250],
	124:  _DwarfOp_name[1250:1262],
	125:  _DwarfOp_name[1262:1274],
	126:  _DwarfOp_name[1274:1286],
	127:  _DwarfOp_name[1286:1298],
	128:  _DwarfOp_name[1298:1310],
	129:  _DwarfOp_name[1310:1322],
	130:  _DwarfOp_name[1322:1334],
	131:  _DwarfOp_name[1334:1346],
	132:  _DwarfOp_name[1346:1358],
	133:  _DwarfOp_name[1358:1370],
	134:  _DwarfOp_name[1370:1382],
	135:  _DwarfOp_name[1382:1394],
	136:  _DwarfOp_name[1394:1406],
	137:  _DwarfOp_name[1406:1418],
	138:  _DwarfOp_name[1418:1430],
	139:  _DwarfOp_name[1430:1442],
	140:  _DwarfOp_name[1442:1454],
	141:  _DwarfOp_name[1454:1466],
	142:  _DwarfOp_name[1466:1478],
	143:  _DwarfOp_name[1478:1490],
	144:  _DwarfOp_name[1490:1500],
	145:  _DwarfOp_name[1500:1511],
	146:  _DwarfOp_name[1511:1522],
	147:  _DwarfOp_name[1522:1533],
	148:  _DwarfOp_name[1533:1549],
	149:  _DwarfOp_name[1549:1566],
	150:  _DwarfOp_name[1566:1575],
	151:  _DwarfOp_name[1575:1600],
	152:  _DwarfOp_name[1600:1611],
	153:  _DwarfOp_name[1611:1622],
	154:  _DwarfOp_name[1622:1636],
	155:  _DwarfOp_name[1636:1658],
	156:  _DwarfOp_name[1658:1678],
	157:  _DwarfOp_name[1678:1693],
	158:  _DwarfOp_name[1693:1713],
	159:  _DwarfOp_name[1713:1730],
	160:  _DwarfOp_name[1730:1752],
	161:  _DwarfOp_name[1752:1763],
	162:  _DwarfOp_name[1763:1775],
	163:  _DwarfOp_name[1775:1792],
	164:  _DwarfOp_name[1792:1808],
	165:  _DwarfOp_name[1808:1825],
	166:  _DwarfOp_name[1825:1841],
	167:  _DwarfOp_name[1841:1858],
	168:  _DwarfOp_name[1858:1871],
	169:  _DwarfOp_name[1871:1888],
	224:  _DwarfOp_name[1888:1914],
	225:  _DwarfOp_name[1914:1931],
	226:  _DwarfOp_name[1931:1949],
	227:  _DwarfOp_name[1949:1967],
	228:  _DwarfOp_name[1967:1985],
	229:  _DwarfOp_name[1985:2005],
	230:  _DwarfOp_name[2005:2017],
	232:  _DwarfOp_name[2017:2038],
	237:  _DwarfOp_name[2038:2057],
	238:  _DwarfOp_name[2057:2080],
	240:  _DwarfOp_name[2080:2098],
	243:  _DwarfOp_name[2098:2119],
	248:  _DwarfOp_name[2119:2143],
	251:  _DwarfOp_name[2143:2163],
	252:  _DwarfOp_name[2163:2184],
	4096: _DwarfOp_name[2184:2203],
	4097: _DwarfOp_name[2203:2221],
	4098: _DwarfOp_name[2221:2242],
	4099: _DwarfOp_name[2242:2264],
	4100: _DwarfOp_name[2264:2291],
	4101: _DwarfOp_name[2291:2305],
}

func (i DwarfOp) String() string {
	if str, ok := _DwarfOp_map[i]; ok {
		return str
	}
	return "DwarfOp(" + strconv.FormatInt(int64(i), 10) + ")"
}
