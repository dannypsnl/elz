// Code generated from Elz.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 57, 337,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 6, 50, 266, 10, 50, 13, 50,
	14, 50, 267, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 7, 51, 276, 10,
	51, 12, 51, 14, 51, 279, 11, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3,
	52, 7, 52, 287, 10, 52, 12, 52, 14, 52, 290, 11, 52, 3, 53, 5, 53, 293,
	10, 53, 3, 54, 3, 54, 5, 54, 297, 10, 54, 3, 55, 7, 55, 300, 10, 55, 12,
	55, 14, 55, 303, 11, 55, 3, 55, 3, 55, 6, 55, 307, 10, 55, 13, 55, 14,
	55, 308, 3, 56, 6, 56, 312, 10, 56, 13, 56, 14, 56, 313, 3, 57, 3, 57,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 327,
	10, 58, 3, 59, 3, 59, 7, 59, 331, 10, 59, 12, 59, 14, 59, 334, 11, 59,
	3, 59, 3, 59, 4, 277, 332, 2, 60, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53,
	105, 2, 107, 2, 109, 54, 111, 55, 113, 2, 115, 56, 117, 57, 3, 2, 6, 5,
	2, 11, 12, 15, 15, 34, 34, 16, 2, 67, 92, 97, 97, 99, 124, 194, 216, 218,
	248, 250, 769, 882, 895, 897, 8193, 8206, 8207, 8306, 8593, 11266, 12273,
	12291, 55297, 63746, 64977, 65010, 65535, 6, 2, 50, 59, 185, 185, 770,
	881, 8257, 8258, 3, 2, 50, 59, 2, 342, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 109, 3,
	2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 3,
	119, 3, 2, 2, 2, 5, 122, 3, 2, 2, 2, 7, 124, 3, 2, 2, 2, 9, 127, 3, 2,
	2, 2, 11, 130, 3, 2, 2, 2, 13, 133, 3, 2, 2, 2, 15, 136, 3, 2, 2, 2, 17,
	138, 3, 2, 2, 2, 19, 140, 3, 2, 2, 2, 21, 142, 3, 2, 2, 2, 23, 144, 3,
	2, 2, 2, 25, 146, 3, 2, 2, 2, 27, 148, 3, 2, 2, 2, 29, 151, 3, 2, 2, 2,
	31, 158, 3, 2, 2, 2, 33, 165, 3, 2, 2, 2, 35, 170, 3, 2, 2, 2, 37, 172,
	3, 2, 2, 2, 39, 174, 3, 2, 2, 2, 41, 180, 3, 2, 2, 2, 43, 183, 3, 2, 2,
	2, 45, 185, 3, 2, 2, 2, 47, 187, 3, 2, 2, 2, 49, 189, 3, 2, 2, 2, 51, 193,
	3, 2, 2, 2, 53, 195, 3, 2, 2, 2, 55, 197, 3, 2, 2, 2, 57, 200, 3, 2, 2,
	2, 59, 205, 3, 2, 2, 2, 61, 207, 3, 2, 2, 2, 63, 209, 3, 2, 2, 2, 65, 213,
	3, 2, 2, 2, 67, 217, 3, 2, 2, 2, 69, 220, 3, 2, 2, 2, 71, 225, 3, 2, 2,
	2, 73, 231, 3, 2, 2, 2, 75, 233, 3, 2, 2, 2, 77, 235, 3, 2, 2, 2, 79, 238,
	3, 2, 2, 2, 81, 240, 3, 2, 2, 2, 83, 242, 3, 2, 2, 2, 85, 244, 3, 2, 2,
	2, 87, 247, 3, 2, 2, 2, 89, 250, 3, 2, 2, 2, 91, 253, 3, 2, 2, 2, 93, 256,
	3, 2, 2, 2, 95, 259, 3, 2, 2, 2, 97, 262, 3, 2, 2, 2, 99, 265, 3, 2, 2,
	2, 101, 271, 3, 2, 2, 2, 103, 284, 3, 2, 2, 2, 105, 292, 3, 2, 2, 2, 107,
	296, 3, 2, 2, 2, 109, 301, 3, 2, 2, 2, 111, 311, 3, 2, 2, 2, 113, 315,
	3, 2, 2, 2, 115, 326, 3, 2, 2, 2, 117, 328, 3, 2, 2, 2, 119, 120, 7, 41,
	2, 2, 120, 121, 7, 107, 2, 2, 121, 4, 3, 2, 2, 2, 122, 123, 7, 58, 2, 2,
	123, 6, 3, 2, 2, 2, 124, 125, 7, 51, 2, 2, 125, 126, 7, 56, 2, 2, 126,
	8, 3, 2, 2, 2, 127, 128, 7, 53, 2, 2, 128, 129, 7, 52, 2, 2, 129, 10, 3,
	2, 2, 2, 130, 131, 7, 56, 2, 2, 131, 132, 7, 54, 2, 2, 132, 12, 3, 2, 2,
	2, 133, 134, 7, 41, 2, 2, 134, 135, 7, 104, 2, 2, 135, 14, 3, 2, 2, 2,
	136, 137, 7, 37, 2, 2, 137, 16, 3, 2, 2, 2, 138, 139, 7, 93, 2, 2, 139,
	18, 3, 2, 2, 2, 140, 141, 7, 95, 2, 2, 141, 20, 3, 2, 2, 2, 142, 143, 7,
	42, 2, 2, 143, 22, 3, 2, 2, 2, 144, 145, 7, 46, 2, 2, 145, 24, 3, 2, 2,
	2, 146, 147, 7, 43, 2, 2, 147, 26, 3, 2, 2, 2, 148, 149, 7, 60, 2, 2, 149,
	150, 7, 60, 2, 2, 150, 28, 3, 2, 2, 2, 151, 152, 7, 107, 2, 2, 152, 153,
	7, 111, 2, 2, 153, 154, 7, 114, 2, 2, 154, 155, 7, 113, 2, 2, 155, 156,
	7, 116, 2, 2, 156, 157, 7, 118, 2, 2, 157, 30, 3, 2, 2, 2, 158, 159, 7,
	116, 2, 2, 159, 160, 7, 103, 2, 2, 160, 161, 7, 118, 2, 2, 161, 162, 7,
	119, 2, 2, 162, 163, 7, 116, 2, 2, 163, 164, 7, 112, 2, 2, 164, 32, 3,
	2, 2, 2, 165, 166, 7, 110, 2, 2, 166, 167, 7, 113, 2, 2, 167, 168, 7, 113,
	2, 2, 168, 169, 7, 114, 2, 2, 169, 34, 3, 2, 2, 2, 170, 171, 7, 125, 2,
	2, 171, 36, 3, 2, 2, 2, 172, 173, 7, 127, 2, 2, 173, 38, 3, 2, 2, 2, 174,
	175, 7, 111, 2, 2, 175, 176, 7, 99, 2, 2, 176, 177, 7, 118, 2, 2, 177,
	178, 7, 101, 2, 2, 178, 179, 7, 106, 2, 2, 179, 40, 3, 2, 2, 2, 180, 181,
	7, 63, 2, 2, 181, 182, 7, 64, 2, 2, 182, 42, 3, 2, 2, 2, 183, 184, 7, 63,
	2, 2, 184, 44, 3, 2, 2, 2, 185, 186, 7, 62, 2, 2, 186, 46, 3, 2, 2, 2,
	187, 188, 7, 64, 2, 2, 188, 48, 3, 2, 2, 2, 189, 190, 7, 48, 2, 2, 190,
	191, 7, 48, 2, 2, 191, 192, 7, 48, 2, 2, 192, 50, 3, 2, 2, 2, 193, 194,
	7, 61, 2, 2, 194, 52, 3, 2, 2, 2, 195, 196, 7, 66, 2, 2, 196, 54, 3, 2,
	2, 2, 197, 198, 7, 47, 2, 2, 198, 199, 7, 64, 2, 2, 199, 56, 3, 2, 2, 2,
	200, 201, 7, 107, 2, 2, 201, 202, 7, 111, 2, 2, 202, 203, 7, 114, 2, 2,
	203, 204, 7, 110, 2, 2, 204, 58, 3, 2, 2, 2, 205, 206, 7, 60, 2, 2, 206,
	60, 3, 2, 2, 2, 207, 208, 7, 45, 2, 2, 208, 62, 3, 2, 2, 2, 209, 210, 7,
	110, 2, 2, 210, 211, 7, 103, 2, 2, 211, 212, 7, 118, 2, 2, 212, 64, 3,
	2, 2, 2, 213, 214, 7, 111, 2, 2, 214, 215, 7, 119, 2, 2, 215, 216, 7, 118,
	2, 2, 216, 66, 3, 2, 2, 2, 217, 218, 7, 104, 2, 2, 218, 219, 7, 112, 2,
	2, 219, 68, 3, 2, 2, 2, 220, 221, 7, 118, 2, 2, 221, 222, 7, 123, 2, 2,
	222, 223, 7, 114, 2, 2, 223, 224, 7, 103, 2, 2, 224, 70, 3, 2, 2, 2, 225,
	226, 7, 118, 2, 2, 226, 227, 7, 116, 2, 2, 227, 228, 7, 99, 2, 2, 228,
	229, 7, 107, 2, 2, 229, 230, 7, 118, 2, 2, 230, 72, 3, 2, 2, 2, 231, 232,
	7, 40, 2, 2, 232, 74, 3, 2, 2, 2, 233, 234, 7, 44, 2, 2, 234, 76, 3, 2,
	2, 2, 235, 236, 7, 99, 2, 2, 236, 237, 7, 117, 2, 2, 237, 78, 3, 2, 2,
	2, 238, 239, 7, 96, 2, 2, 239, 80, 3, 2, 2, 2, 240, 241, 7, 49, 2, 2, 241,
	82, 3, 2, 2, 2, 242, 243, 7, 47, 2, 2, 243, 84, 3, 2, 2, 2, 244, 245, 7,
	62, 2, 2, 245, 246, 7, 63, 2, 2, 246, 86, 3, 2, 2, 2, 247, 248, 7, 64,
	2, 2, 248, 249, 7, 63, 2, 2, 249, 88, 3, 2, 2, 2, 250, 251, 7, 35, 2, 2,
	251, 252, 7, 63, 2, 2, 252, 90, 3, 2, 2, 2, 253, 254, 7, 63, 2, 2, 254,
	255, 7, 63, 2, 2, 255, 92, 3, 2, 2, 2, 256, 257, 7, 40, 2, 2, 257, 258,
	7, 40, 2, 2, 258, 94, 3, 2, 2, 2, 259, 260, 7, 126, 2, 2, 260, 261, 7,
	126, 2, 2, 261, 96, 3, 2, 2, 2, 262, 263, 7, 65, 2, 2, 263, 98, 3, 2, 2,
	2, 264, 266, 9, 2, 2, 2, 265, 264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267,
	265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270,
	8, 50, 2, 2, 270, 100, 3, 2, 2, 2, 271, 272, 7, 49, 2, 2, 272, 273, 7,
	49, 2, 2, 273, 277, 3, 2, 2, 2, 274, 276, 11, 2, 2, 2, 275, 274, 3, 2,
	2, 2, 276, 279, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2,
	278, 280, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 280, 281, 7, 12, 2, 2, 281,
	282, 3, 2, 2, 2, 282, 283, 8, 51, 2, 2, 283, 102, 3, 2, 2, 2, 284, 288,
	5, 105, 53, 2, 285, 287, 5, 107, 54, 2, 286, 285, 3, 2, 2, 2, 287, 290,
	3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 104, 3, 2,
	2, 2, 290, 288, 3, 2, 2, 2, 291, 293, 9, 3, 2, 2, 292, 291, 3, 2, 2, 2,
	293, 106, 3, 2, 2, 2, 294, 297, 5, 105, 53, 2, 295, 297, 9, 4, 2, 2, 296,
	294, 3, 2, 2, 2, 296, 295, 3, 2, 2, 2, 297, 108, 3, 2, 2, 2, 298, 300,
	5, 113, 57, 2, 299, 298, 3, 2, 2, 2, 300, 303, 3, 2, 2, 2, 301, 299, 3,
	2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 304, 3, 2, 2, 2, 303, 301, 3, 2, 2,
	2, 304, 306, 7, 48, 2, 2, 305, 307, 5, 113, 57, 2, 306, 305, 3, 2, 2, 2,
	307, 308, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309,
	110, 3, 2, 2, 2, 310, 312, 5, 113, 57, 2, 311, 310, 3, 2, 2, 2, 312, 313,
	3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 112, 3, 2,
	2, 2, 315, 316, 9, 5, 2, 2, 316, 114, 3, 2, 2, 2, 317, 318, 7, 118, 2,
	2, 318, 319, 7, 116, 2, 2, 319, 320, 7, 119, 2, 2, 320, 327, 7, 103, 2,
	2, 321, 322, 7, 104, 2, 2, 322, 323, 7, 99, 2, 2, 323, 324, 7, 110, 2,
	2, 324, 325, 7, 117, 2, 2, 325, 327, 7, 103, 2, 2, 326, 317, 3, 2, 2, 2,
	326, 321, 3, 2, 2, 2, 327, 116, 3, 2, 2, 2, 328, 332, 7, 36, 2, 2, 329,
	331, 11, 2, 2, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 333,
	3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334, 332, 3, 2,
	2, 2, 335, 336, 7, 36, 2, 2, 336, 118, 3, 2, 2, 2, 13, 2, 267, 277, 288,
	292, 296, 301, 308, 313, 326, 332, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "''i'", "'8'", "'16'", "'32'", "'64'", "''f'", "'#'", "'['", "']'",
	"'('", "','", "')'", "'::'", "'import'", "'return'", "'loop'", "'{'", "'}'",
	"'match'", "'=>'", "'='", "'<'", "'>'", "'...'", "';'", "'@'", "'->'",
	"'impl'", "':'", "'+'", "'let'", "'mut'", "'fn'", "'type'", "'trait'",
	"'&'", "'*'", "'as'", "'^'", "'/'", "'-'", "'<='", "'>='", "'!='", "'=='",
	"'&&'", "'||'", "'?'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "COMMENT", "ID",
	"FLOAT", "INT", "BOOLEAN", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "WS", "COMMENT",
	"ID", "StartLetter", "Letter", "FLOAT", "INT", "Digit", "BOOLEAN", "STRING",
}

type ElzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewElzLexer(input antlr.CharStream) *ElzLexer {

	l := new(ElzLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Elz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ElzLexer tokens.
const (
	ElzLexerT__0    = 1
	ElzLexerT__1    = 2
	ElzLexerT__2    = 3
	ElzLexerT__3    = 4
	ElzLexerT__4    = 5
	ElzLexerT__5    = 6
	ElzLexerT__6    = 7
	ElzLexerT__7    = 8
	ElzLexerT__8    = 9
	ElzLexerT__9    = 10
	ElzLexerT__10   = 11
	ElzLexerT__11   = 12
	ElzLexerT__12   = 13
	ElzLexerT__13   = 14
	ElzLexerT__14   = 15
	ElzLexerT__15   = 16
	ElzLexerT__16   = 17
	ElzLexerT__17   = 18
	ElzLexerT__18   = 19
	ElzLexerT__19   = 20
	ElzLexerT__20   = 21
	ElzLexerT__21   = 22
	ElzLexerT__22   = 23
	ElzLexerT__23   = 24
	ElzLexerT__24   = 25
	ElzLexerT__25   = 26
	ElzLexerT__26   = 27
	ElzLexerT__27   = 28
	ElzLexerT__28   = 29
	ElzLexerT__29   = 30
	ElzLexerT__30   = 31
	ElzLexerT__31   = 32
	ElzLexerT__32   = 33
	ElzLexerT__33   = 34
	ElzLexerT__34   = 35
	ElzLexerT__35   = 36
	ElzLexerT__36   = 37
	ElzLexerT__37   = 38
	ElzLexerT__38   = 39
	ElzLexerT__39   = 40
	ElzLexerT__40   = 41
	ElzLexerT__41   = 42
	ElzLexerT__42   = 43
	ElzLexerT__43   = 44
	ElzLexerT__44   = 45
	ElzLexerT__45   = 46
	ElzLexerT__46   = 47
	ElzLexerT__47   = 48
	ElzLexerWS      = 49
	ElzLexerCOMMENT = 50
	ElzLexerID      = 51
	ElzLexerFLOAT   = 52
	ElzLexerINT     = 53
	ElzLexerBOOLEAN = 54
	ElzLexerSTRING  = 55
)
