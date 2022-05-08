package bios

// #include <stdio.h>
//
// typedef void (*EMPTY_FNC)(void);
// typedef int16_t (*GBA_SQRT_FNC)(int32_t);
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
// int16_t gba_sqrt(int32_t r0) {
//    // 0x08	Sqrt
//    GBA_SQRT_FNC f = (GBA_SQRT_FNC)0x08;
//    return f(r0);
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

// Function descriptions taken from here: https://web.archive.org/web/20090223235549/http://nocash.emubase.de/gbatek.htm#biosfunctions

// SoftReset clears 200h bytes of RAM (containing stacks, and BIOS IRQ vector/flags), initializes system,
// supervisor, and irq stack pointers, sets R0-R12, LR_svc, SPSR_svc, LR_irq, and SPSR_irq to zero, and
// enters system mode.
// Return: Does not return to calling procedure, instead, loads the above return address into R14, and then
// jumps to that address by a "BX R14" opcode.
func SoftReset() {
	C.gba_soft_reset()
}

// RegisterRamReset resets the I/O registers and RAM specified in ResetFlags. However, it does not clear the
// CPU internal RAM area from 3007E00h-3007FFFh.
// Return: No return value.
// The function always switches the screen into forced blank by setting DISPCNT=0080h (regardless of incoming R0, screen becomes white).
func RegisterRamReset() {
	C.gba_register_ram_reset()
}

// Halt halts the CPU until an interrupt request occurs. The CPU is switched into low-power mode,
// all other circuits (video, sound, timers, serial, keypad, system clock) are kept operating.
// Halt mode is terminated when any enabled interrupts are requested, that is when (IE AND IF) is
// not zero, the GBA locks up if that condition doesn't get true. However, the state of CPUs IRQ
// disable bit in CPSR register, and the IME register are don't care, Halt passes through even if
// either one has disabled interrupts.
// On GBA and NDS7, Halt is implemented by writing to HALTCNT, Port 4000301h.
// No parameters, no return value.
func Halt() {
	C.gba_halt()
}

// Stop switches the GBA into very low power mode (to be used similar as a screen-saver). The CPU,
// System Clock, Sound, Video, SIO-Shift Clock, DMAs, and Timers are stopped.
// Stop state can be terminated by the following interrupts only (as far as enabled in IE
// register): Joypad, Game Pak, or General-Purpose-SIO.
// "The system clock is stopped so the IF flag is not set."
// Preparation for Stop:
// Disable Video before implementing Stop (otherwise Video just freezes, but still keeps consuming
// battery power). Possibly required to disable Sound also? Obviously, it'd be also recommended to
// disable any external hardware (such like Rumble or Infra-Red) as far as possible.
// No parameters, no return value.
func Stop() {
	C.gba_stop()
}

// IntrWait continues to wait in Halt state until one (or more) of the specified interrupt(s) do occur.
// The function forcefully sets IME=1. When using multiple interrupts at the same time, this function is
// having less overhead than repeatedly calling the Halt function.
// Caution: When using IntrWait or VBlankIntrWait, the user interrupt handler MUST update the BIOS Interrupt
// Flags value in RAM; when acknowleding processed interrupt(s) by writing a value to the IF register, the
// same value should be also ORed to the BIOS Interrupt Flags value, at following memory location:
//
//   Host     GBA (16bit)  NDS7 (32bit)  NDS9 (32bit)
//   Address  [3007FF8h]   [380FFF8h]    [DTCM+3FF8h]
//
// Return: No return value, the selected flag(s) are automatically reset in BIOS Interrupt Flags value in RAM upon return.
func IntrWait() {
	C.gba_intr_wait()
}

// VBlankIntrWait continues to wait in Halt status until a new V-Blank interrupt occurs.
// The function sets r0=1 and r1=1 and then executes IntrWait (SWI 04h), see IntrWait for details.
// No parameters, no return value.
func VBlankIntrWait() {
	C.gba_v_blank_intr_wait()
}

// Div signed Division, r0/r1 (SWI 06h).
// For example, incoming -1234, 10 should return -123, -4, +123.
// The function usually gets caught in an endless loop upon division by zero.
func Div() {
	C.gba_div()
}

// Div signed Division, but incoming parameters are exchanged, r1/r0 (SWI 07h). Slightly slower (3 clock cycles) than SWI 06h.
func DivArm() {
	C.gba_div_arm()
}

// Sqrt calculate square root.
// The result is an integer value, so Sqrt(2) would return 1, to avoid this inaccuracy,
// shift left incoming number by 2*N as much as possible (the result is then shifted left by 1*N).
// Ie. Sqrt(2 shl 30) would return 1.41421 shl 15.
func Sqrt(r0 C.int32_t) C.int16_t {
	return C.gba_sqrt(r0)
}

// ArcTan calculates the arc tangent.
// Note: there is a problem in accuracy with "THETA<-PI/4, PI/4<THETA".
func ArcTan() {
	C.gba_arc_tan()
}

// ArcTan2 calculates the arc tangent after correction processing.
func ArcTan2() {
	C.gba_arc_tan2()
}

// CPUSet Memory copy/fill in units of 4 bytes or 2 bytes. Memcopy is implemented as
// repeated LDMIA/STMIA [Rb]!,r3 or LDRH/STRH r3,[r0,r5] instructions. Memfill as single LDMIA or
// LDRH followed by repeated STMIA [Rb]!,r3 or STRH r3,[r0,r5].
// The length must be a multiple of 4 bytes (32bit mode) or 2 bytes (16bit mode). The (half)wordcount
// in r2 must be length/4 (32bit mode) or length/2 (16bit mode), ie. length in word/halfword units rather than byte units.
// Return: No return value, Data written to destination address.
func CPUSet() {
	C.gba_cpu_set()
}

// CPUFastSet memory copy/fill in units of 32 bytes. Memcopy is implemented as repeated LDMIA/STMIA [Rb]!,r2-r9 instructions.
// Memfill as single LDR followed by repeated STMIA [Rb]!,r2-r9. After processing all 32-byte-blocks, the NDS additonally
// processes the remaining words as 4-byte blocks.
// The length is specifed as wordcount, ie. the number of bytes divided by 4.
// On the GBA, the length must be a multiple of 8 words (32 bytes).
// Return: No return value, Data written to destination address.
func CPUFastSet() {
	C.gba_cpu_fast_set()
}

// BiosChecksum calculates the checksum of the BIOS ROM (by reading in 32bit units, and adding up these values).
// IRQ and FIQ are disabled during execution.
// The checksum is BAAE187Fh (GBA and GBA SP), or BAAE1880h (DS in GBA mode, whereas the only difference is that
// the byte at [3F0Ch] is changed from 00h to 01h, otherwise the BIOS is 1:1 same as GBA BIOS, it does even include multiboot code).
// Parameters: None. Return: r0=Checksum.
func BiosChecksum() {
	C.gba_bios_checksum()
}

// BgAffineSet used to calculate BG Rotation/Scaling parameters.
// Return: No return value, Data written to destination address.
func BgAffineSet() {
	C.gba_bg_affine_set()
}

// ObjAffineSet calculates and sets the OBJ's affine parameters from the scaling ratio and angle of rotation.
// The affine parameters are calculated from the parameters set in Srcp.
// The four affine parameters are set every Offset bytes, starting from the Destp address.
// If the Offset value is 2, the parameters are stored contiguously. If the value is 8, they match the structure of OAM.
// When Srcp is arrayed, the calculation can be performed continuously by specifying Num.
// Return: No return value, Data written to destination address.
func ObjAffineSet() {
	C.gba_obj_affine_set()
}

// BitUnPack used to increase the color depth of bitmaps or tile data.
// For example, to convert a 1bit monochrome font into 4bit or 8bit GBA tiles.
// The Unpack Info is specified separately, allowing to convert the same source data into different formats.
// Data is written in 32bit units, Destination can be Wram or Vram. The size of unpacked data must be a multiple of 4 bytes.
// The width of source units (plus the offset) should not exceed the destination width.
// Return: No return value, Data written to destination address.
func BitUnPack() {
	C.gba_bit_un_pack()
}
func LZ77UnCompWRAM() {
	C.gba_lz77_un_comp_wram()
}
func LZ77UnCompVRAM() {
	C.gba_lz77_un_comp_vram()
}

// HuffUnComp the decoder starts in root node, the separate bits in the bitstream specify if the next node is node0
// or node1, if that node is a data node, then the data is stored in memory, and the decoder is reset to the root node.
// The most often used data should be as close to the root node as possible. For example, the 4-byte string "Huff" could
// be compressed to 6 bits: 10-11-0-0, with root.0 pointing directly to data "f", and root.1 pointing to a child node,
// whose nodes point to data "H" and data "u".
// Data is written in units of 32bits, if the size of the compressed data is not a multiple of 4, please adjust it as much
// as possible by padding with 0.
// Align the source address to a 4Byte boundary.
// Return: No return value, Data written to destination address.
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

// SoundBiasChange increments or decrements the current level of the SOUNDBIAS register (with short delays) until reaching
// the desired new level. The upper bits of the register are kept unchanged.
// Return: No return value.
func SoundBiasChange() {
	C.gba_sound_bias_change()
}

// SoundDriverInit initializes the sound driver. Call this only once when the game starts up.
// It is essential that the work area already be secured at the time this function is called.
// You cannot execute this driver multiple times, even if separate work areas have been prepared.
// r0  Pointer to work area for sound driver, SoundArea structure as follows:
//
// SoundArea (sa) Structure
//  u32    ident      Flag the system checks to see whether the
// 				   work area has been initialized and whether it
// 				   is currently being accessed.
//  vu8    DmaCount   User access prohibited
//  u8     reverb     Variable for applying reverb effects to direct sound
//  u16    d1         User access prohibited
//  void   (*func)()  User access prohibited
//  int    intp       User access prohibited
//  void*  NoUse      User access prohibited
//  SndCh  vchn[MAX]  The structure array for controlling the direct
// 				   sound channels (currently 8 channels are
// 				   available). The term "channel" here does
// 				   not refer to hardware channels, but rather to
// 				   virtual constructs inside the sound driver.
//  s8     pcmbuf[PCM_BF*2]
// SoundChannel Structure
//  u8         sf     The flag indicating the status of this channel.
// 				   When 0 sound is stopped.
// 				   To start sound, set other parameters and
// 				   then write 80h to here.
// 				   To stop sound, logical OR 40h for a
// 				   release-attached off (key-off), or write zero
// 				   for a pause. The use of other bits is
// 				   prohibited.
//  u8         r1     User access prohibited
//  u8         rv     Sound volume output to right side
//  u8         lv     Sound volume output to left side
//  u8         at     The attack value of the envelope. When the
// 				   sound starts, the volume begins at zero and
// 				   increases every 1/60 second. When it
// 				   reaches 255, the process moves on to the
// 				   next decay value.
//  u8         de     The decay value of the envelope. It is
// 				   multiplied by "this value/256" every 1/60
// 				   sec. and when sustain value is reached, the
// 				   process moves to the sustain condition.
//  u8         su     The sustain value of the envelope. The
// 				   sound is sustained by this amount.
// 				   (Actually, multiplied by rv/256, lv/256 and
// 				   output left and right.)
//  u8         re     The release value of the envelope. Key-off
// 				   (logical OR 40h in sf) to enter this state.
// 				   The value is multiplied by "this value/256"
// 				   every 1/60 sec. and when it reaches zero,
// 				   this channel is completely stopped.
//  u8         r2[4]  User access prohibited
//  u32        fr     The frequency of the produced sound.
// 				   Write the value obtained with the
// 				   MidiKey2Freq function here.
//  WaveData*  wp     Pointer to the sound's waveform data. The waveform
// 				   data can be generated automatically from the AIFF
// 				   file using the tool (aif2agb.exe), so users normally
// 				   do not need to create this themselves.
//  u32        r3[6]  User access prohibited
//  u8         r4[4]  User access prohibited
// WaveData Structure
//  u16   type    Indicates the data type. This is currently not used.
//  u16   stat    At the present time, non-looped (1 shot) waveform
// 			   is 0000h and forward loop is 4000h.
//  u32   freq    This value is used to calculate the frequency.
// 			   It is obtained using the following formula:
// 			   sampling rate x 2^((180-original MIDI key)/12)
//  u32   loop    Loop pointer (start of loop)
//  u32   size    Number of samples (end position)
//  s8    data[]  The actual waveform data. Takes (number of samples+1)
// 			   bytes of 8bit signed linear uncompressed data. The last
// 			   byte is zero for a non-looped waveform, and the same
// 			   value as the loop pointer data for a looped waveform.
//
// Return: No return value.
func SoundDriverInit() {
	C.gba_sound_driver_init()
}

// SoundDriverMode sets the sound driver operation mode.
//   r0  Sound driver operation mode
//   Bit    Expl.
//   0-6    Direct Sound Reverb value (0-127, default=0) (ignored if Bit7=0)
//   7      Direct Sound Reverb set (0=ignore, 1=apply reverb value)
//   8-11   Direct Sound Simultaneously-produced (1-12 channels, default 8)
//   12-15  Direct Sound Master volume (1-15, default 15)
//   16-19  Direct Sound Playback Frequency (1-12 = 5734,7884,10512,13379,
// 		 15768,18157,21024,26758,31536,36314,40137,42048, def 4=13379 Hz)
//   20-23  Final number of D/A converter bits (8-11 = 9-6bits, def. 9=8bits)
//   24-31  Not used.
// Return: No return value.
func SoundDriverMode() {
	C.gba_sound_driver_mode()
}

// SoundDriverMain main of the sound driver.
// Call every 1/60 of a second. The flow of the process is to call SoundDriverVSync, which is explained later,
// immediately after the V-Blank interrupt.
// After that, this routine is called after BG and OBJ processing is executed.
// No parameters, no return value.
func SoundDriverMain() {
	C.gba_sound_driver_main()
}

// SoundDriverVSync an extremely short system call that resets the sound DMA. The timing is extremely critical,
// so call this function immediately after the V-Blank interrupt every 1/60 second.
// No parameters, no return value.
func SoundDriverVSync() {
	C.gba_sound_driver_v_sync()
}

// SoundChannelClear clears all direct sound channels and stops the sound.
// This function may not operate properly when the library which expands the sound driver feature is combined afterwards.
// In this case, do not use it.
// No parameters, no return value.
func SoundChannelClear() {
	C.gba_sound_channel_clear()
}

// MidiKey2Freq calculates the value of the assignment to ((SoundArea)sa).vchn[x].fr when playing the wave
// data, wa, with the interval (MIDI KEY) mk and the fine adjustment value (halftones=256) fp.
//
//  r0  WaveData* wa
//  r1  u8 mk
//  r2  u8 fp
//
// Return:
//
//  r0  u32
//
// This function is particularly popular because it allows to read from BIOS memory without copy protection
// range checks. The formula to read one byte (a) from address (i, 0..3FFF) is:
// a = (MidiKey2Freq(i-(((i AND 3)+1)OR 3), 168, 0) * 2) SHR 24
func MIDIKey2Freq() {
	C.gba_midi_key2_freq()
}

// SoundDriverVSyncOff due to problems with the main program if the V-Blank interrupts are stopped, and
// SoundDriverVSync cannot be called every 1/60 a second, this function must be used to stop sound DMA.
// Otherwise, even if you exceed the limit of the buffer the DMA will not stop and noise will result.
// No parameters, no return value.
func SoundDriverVSyncOff() {
	C.gba_sound_driver_v_sync_off()
}

// SoundDriverVSyncOn this function restarts the sound DMA stopped with the previously described SoundDriverVSyncOff.
// After calling this function, have a V-Blank occur within 2/60 of a second and call SoundDriverVSync.
// No parameters, no return value.
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

// MultiBoot this function uploads & starts program code to slave GBAs, allowing to launch programs on
// slave units even if no cartridge is inserted into the slaves (this works because all GBA BIOSes contain
// built-in download procedures in ROM).
// However, the SWI 25h BIOS upload function covers only 45% of the required Transmission Protocol, the other
// 55% must be coded in the master cartridge (see Transmission Protocol below).
//
// r0  Pointer to MultiBootParam structure
// r1  Transfer Mode (undocumented)
//     0=256KHz, 32bit, Normal mode    (fast and stable)
//     1=115KHz, 16bit, MultiPlay mode (default, slow, up to three slaves)
//     2=2MHz,   32bit, Normal mode    (fastest but maybe unstable)
//
// Note: HLL-programmers that are using the MultiBoot(param_ptr) macro cannot
// specify the transfer mode and will be forcefully using MultiPlay mode.
// Return:
// r0  0=okay, 1=failed
func MultiBoot() {
	C.gba_multi_boot()
}

// HardReset this function reboots the GBA (including for getting through the time-consuming nintendo intro,
// which is making the function particularly useless and annoying).
// Parameters: None. Return: Never/Reboot.
// Execution Time: About 2 seconds (!)
func HardReset() {
	C.gba_hard_reset()
}

// CustomHalt writes the 8bit parameter value to HALTCNT, below values are equivalent to Halt and Stop/Sleep functions,
// other values reserved, purpose unknown.
//
//   r2  8bit parameter (GBA: 00h=Halt, 80h=Stop) (NDS7: 80h=Halt, C0h=Sleep)
//
// No return value.
func CustomHalt() {
	C.gba_custom_halt()
}
