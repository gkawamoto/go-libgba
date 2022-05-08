package bios

// typedef void (*EMPTY_FNC)(void);
//
// void gba_soft_reset() {
//    // 0x00	SoftReset
//    EMPTY_FNC f = (EMPTY_FNC)0x00;
//    f();
// }
// void gba_register_ram_reset() {
//    // 0x01	RegisterRamReset
//    EMPTY_FNC f = (EMPTY_FNC)0x01;
//    f();
// }
// void gba_halt() {
//    // 0x02	Halt
//    EMPTY_FNC f = (EMPTY_FNC)0x02;
//    f();
// }
// void gba_stop() {
//    // 0x03	Stop
//    EMPTY_FNC f = (EMPTY_FNC)0x03;
//    f();
// }
// void gba_intr_wait() {
//    // 0x04	IntrWait
//    EMPTY_FNC f = (EMPTY_FNC)0x04;
//    f();
// }
// void gba_v_blank_intr_wait() {
//    // 0x05	VBlankIntrWait
//    EMPTY_FNC f = (EMPTY_FNC)0x05;
//    f();
// }
// void gba_div() {
//    // 0x06	Div
//    EMPTY_FNC f = (EMPTY_FNC)0x06;
//    f();
// }
// void gba_div_arm() {
//    // 0x07	DivArm
//    EMPTY_FNC f = (EMPTY_FNC)0x07;
//    f();
// }
// void gba_sqrt() {
//    // 0x08	Sqrt
//    EMPTY_FNC f = (EMPTY_FNC)0x08;
//    f();
// }
// void gba_arc_tan() {
//    // 0x09	ArcTan
//    EMPTY_FNC f = (EMPTY_FNC)0x09;
//    f();
// }
// void gba_arc_tan2() {
//    // 0x0A	ArcTan2
//    EMPTY_FNC f = (EMPTY_FNC)0x0A;
//    f();
// }
// void gba_cpu_set() {
//    // 0x0B	CPUSet
//    EMPTY_FNC f = (EMPTY_FNC)0x0B;
//    f();
// }
// void gba_cpu_fast_set() {
//    // 0x0C	CPUFastSet
//    EMPTY_FNC f = (EMPTY_FNC)0x0C;
//    f();
// }
// void gba_bios_checksum() {
//    // 0x0D	BiosChecksum
//    EMPTY_FNC f = (EMPTY_FNC)0x0D;
//    f();
// }
// void gba_bg_affine_set() {
//    // 0x0E	BgAffineSet
//    EMPTY_FNC f = (EMPTY_FNC)0x0E;
//    f();
// }
// void gba_obj_affine_set() {
//    // 0x0F	ObjAffineSet
//    EMPTY_FNC f = (EMPTY_FNC)0x0F;
//    f();
// }
// void gba_bit_un_pack() {
//    // 0x10	BitUnPack
//    EMPTY_FNC f = (EMPTY_FNC)0x10;
//    f();
// }
// void gba_lz77_un_comp_wram() {
//    // 0x11	LZ77UnCompWRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x11;
//    f();
// }
// void gba_lz77_un_comp_vram() {
//    // 0x12	LZ77UnCompVRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x12;
//    f();
// }
// void gba_huff_un_comp() {
//    // 0x13	HuffUnComp
//    EMPTY_FNC f = (EMPTY_FNC)0x13;
//    f();
// }
// void gba_rl_un_comp_wram() {
//    // 0x14	RLUnCompWRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x14;
//    f();
// }
// void gba_rl_un_comp_vram() {
//    // 0x15	RLUnCompVRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x15;
//    f();
// }
// void gba_diff8bit_un_filter_wram() {
//    // 0x16	Diff8bitUnFilterWRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x16;
//    f();
// }
// void gba_diff8bit_un_filter_vram() {
//    // 0x17	Diff8bitUnFilterVRAM
//    EMPTY_FNC f = (EMPTY_FNC)0x17;
//    f();
// }
// void gba_diff16bit_un_filter() {
//    // 0x18	Diff16bitUnFilter
//    EMPTY_FNC f = (EMPTY_FNC)0x18;
//    f();
// }
// void gba_sound_bias_change() {
//    // 0x19	SoundBiasChange
//    EMPTY_FNC f = (EMPTY_FNC)0x19;
//    f();
// }
// void gba_sound_driver_init() {
//    // 0x1A	SoundDriverInit
//    EMPTY_FNC f = (EMPTY_FNC)0x1A;
//    f();
// }
// void gba_sound_driver_mode() {
//    // 0x1B	SoundDriverMode
//    EMPTY_FNC f = (EMPTY_FNC)0x1B;
//    f();
// }
// void gba_sound_driver_main() {
//    // 0x1C	SoundDriverMain
//    EMPTY_FNC f = (EMPTY_FNC)0x1C;
//    f();
// }
// void gba_sound_driver_v_sync() {
//    // 0x1D	SoundDriverVSync
//    EMPTY_FNC f = (EMPTY_FNC)0x1D;
//    f();
// }
// void gba_sound_channel_clear() {
//    // 0x1E	SoundChannelClear
//    EMPTY_FNC f = (EMPTY_FNC)0x1E;
//    f();
// }
// void gba_midi_key2_freq() {
//    // 0x1F	MIDIKey2Freq
//    EMPTY_FNC f = (EMPTY_FNC)0x1F;
//    f();
// }
// void gba_sound_driver_v_sync_off() {
//    // 0x28	SoundDriverVSyncOff
//    EMPTY_FNC f = (EMPTY_FNC)0x28;
//    f();
// }
// void gba_sound_driver_v_sync_on() {
//    // 0x29	SoundDriverVSyncOn
//    EMPTY_FNC f = (EMPTY_FNC)0x29;
//    f();
// }
// void gba_get_jump_list() {
//    // 0x2A	GetJumpList
//    EMPTY_FNC f = (EMPTY_FNC)0x2A;
//    f();
// }
// void gba_music_player_open() {
//    // 0x20	MusicPlayerOpen
//    EMPTY_FNC f = (EMPTY_FNC)0x20;
//    f();
// }
// void gba_music_player_start() {
//    // 0x21	MusicPlayerStart
//    EMPTY_FNC f = (EMPTY_FNC)0x21;
//    f();
// }
// void gba_music_player_stop() {
//    // 0x22	MusicPlayerStop
//    EMPTY_FNC f = (EMPTY_FNC)0x22;
//    f();
// }
// void gba_music_player_continue() {
//    // 0x23	MusicPlayerContinue
//    EMPTY_FNC f = (EMPTY_FNC)0x23;
//    f();
// }
// void gba_music_player_fade_out() {
//    // 0x24	MusicPlayerFadeOut
//    EMPTY_FNC f = (EMPTY_FNC)0x24;
//    f();
// }
// void gba_multi_boot() {
//    // 0x25	MultiBoot
//    EMPTY_FNC f = (EMPTY_FNC)0x25;
//    f();
// }
// void gba_hard_reset() {
//    // 0x26	HardReset
//    EMPTY_FNC f = (EMPTY_FNC)0x26;
//    f();
// }
// void gba_custom_halt() {
//    // 0x27	CustomHalt
//    EMPTY_FNC f = (EMPTY_FNC)0x27;
//    f();
// }
import "C"

func SoftReset() {
	C.gba_soft_reset()
}
func RegisterRamReset() {
	C.gba_register_ram_reset()
}
func Halt() {
	C.gba_halt()
}
func Stop() {
	C.gba_stop()
}
func IntrWait() {
	C.gba_intr_wait()
}
func VBlankIntrWait() {
	C.gba_v_blank_intr_wait()
}
func Div() {
	C.gba_div()
}
func DivArm() {
	C.gba_div_arm()
}
func Sqrt() {
	C.gba_sqrt()
}
func ArcTan() {
	C.gba_arc_tan()
}
func ArcTan2() {
	C.gba_arc_tan2()
}
func CPUSet() {
	C.gba_cpu_set()
}
func CPUFastSet() {
	C.gba_cpu_fast_set()
}
func BiosChecksum() {
	C.gba_bios_checksum()
}
func BgAffineSet() {
	C.gba_bg_affine_set()
}
func ObjAffineSet() {
	C.gba_obj_affine_set()
}
func BitUnPack() {
	C.gba_bit_un_pack()
}
func LZ77UnCompWRAM() {
	C.gba_lz77_un_comp_wram()
}
func LZ77UnCompVRAM() {
	C.gba_lz77_un_comp_vram()
}
func HuffUnComp() {
	C.gba_huff_un_comp()
}
func RLUnCompWRAM() {
	C.gba_rl_un_comp_wram()
}
func RLUnCompVRAM() {
	C.gba_rl_un_comp_vram()
}
func Diff8bitUnFilterWRAM() {
	C.gba_diff8bit_un_filter_wram()
}
func Diff8bitUnFilterVRAM() {
	C.gba_diff8bit_un_filter_vram()
}
func Diff16bitUnFilter() {
	C.gba_diff16bit_un_filter()
}
func SoundBiasChange() {
	C.gba_sound_bias_change()
}
func SoundDriverInit() {
	C.gba_sound_driver_init()
}
func SoundDriverMode() {
	C.gba_sound_driver_mode()
}
func SoundDriverMain() {
	C.gba_sound_driver_main()
}
func SoundDriverVSync() {
	C.gba_sound_driver_v_sync()
}
func SoundChannelClear() {
	C.gba_sound_channel_clear()
}
func MIDIKey2Freq() {
	C.gba_midi_key2_freq()
}
func SoundDriverVSyncOff() {
	C.gba_sound_driver_v_sync_off()
}
func SoundDriverVSyncOn() {
	C.gba_sound_driver_v_sync_on()
}
func GetJumpList() {
	C.gba_get_jump_list()
}
func MusicPlayerOpen() {
	C.gba_music_player_open()
}
func MusicPlayerStart() {
	C.gba_music_player_start()
}
func MusicPlayerStop() {
	C.gba_music_player_stop()
}
func MusicPlayerContinue() {
	C.gba_music_player_continue()
}
func MusicPlayerFadeOut() {
	C.gba_music_player_fade_out()
}
func MultiBoot() {
	C.gba_multi_boot()
}
func HardReset() {
	C.gba_hard_reset()
}
func CustomHalt() {
	C.gba_custom_halt()
}
