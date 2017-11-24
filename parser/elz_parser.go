// Generated from Elz.g4 by ANTLR 4.7.

package parser // Elz

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 328,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	5, 2, 68, 10, 2, 3, 3, 6, 3, 71, 10, 3, 13, 3, 14, 3, 72, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 81, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 87,
	10, 6, 13, 6, 14, 6, 88, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 104, 10, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 5, 10, 110, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 123, 10, 11, 12, 11, 14, 11, 126,
	11, 11, 3, 11, 5, 11, 129, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 7, 13, 140, 10, 13, 12, 13, 14, 13, 143, 11, 13,
	3, 14, 3, 14, 3, 14, 5, 14, 148, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 157, 10, 16, 12, 16, 14, 16, 160, 11, 16, 3, 17,
	6, 17, 163, 10, 17, 13, 17, 14, 17, 164, 3, 18, 5, 18, 168, 10, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 173, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 178,
	10, 18, 3, 18, 3, 18, 5, 18, 182, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 190, 10, 19, 3, 19, 3, 19, 5, 19, 194, 10, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 5, 21, 201, 10, 21, 3, 21, 3, 21, 3, 21, 5,
	21, 206, 10, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 213, 10, 22,
	3, 22, 3, 22, 3, 22, 7, 22, 218, 10, 22, 12, 22, 14, 22, 221, 11, 22, 3,
	23, 3, 23, 3, 23, 7, 23, 226, 10, 23, 12, 23, 14, 23, 229, 11, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 237, 10, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 242, 10, 25, 3, 25, 3, 25, 3, 25, 5, 25, 247, 10, 25, 3, 25,
	3, 25, 5, 25, 251, 10, 25, 3, 25, 3, 25, 3, 26, 6, 26, 256, 10, 26, 13,
	26, 14, 26, 257, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 266,
	10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 6, 29, 274, 10, 29, 13,
	29, 14, 29, 275, 3, 30, 5, 30, 279, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30,
	284, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 289, 10, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 5, 31, 296, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 306, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 317, 10, 32, 12, 32, 14, 32, 320,
	11, 32, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 326, 10, 33, 3, 33, 2, 3, 62,
	34, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 2, 4, 3, 2, 24,
	25, 4, 2, 17, 17, 26, 26, 2, 344, 2, 67, 3, 2, 2, 2, 4, 70, 3, 2, 2, 2,
	6, 80, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 86, 3, 2, 2, 2, 12, 95, 3, 2,
	2, 2, 14, 97, 3, 2, 2, 2, 16, 100, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20,
	111, 3, 2, 2, 2, 22, 132, 3, 2, 2, 2, 24, 136, 3, 2, 2, 2, 26, 144, 3,
	2, 2, 2, 28, 151, 3, 2, 2, 2, 30, 153, 3, 2, 2, 2, 32, 162, 3, 2, 2, 2,
	34, 167, 3, 2, 2, 2, 36, 185, 3, 2, 2, 2, 38, 197, 3, 2, 2, 2, 40, 200,
	3, 2, 2, 2, 42, 210, 3, 2, 2, 2, 44, 222, 3, 2, 2, 2, 46, 230, 3, 2, 2,
	2, 48, 234, 3, 2, 2, 2, 50, 255, 3, 2, 2, 2, 52, 259, 3, 2, 2, 2, 54, 263,
	3, 2, 2, 2, 56, 273, 3, 2, 2, 2, 58, 278, 3, 2, 2, 2, 60, 290, 3, 2, 2,
	2, 62, 305, 3, 2, 2, 2, 64, 325, 3, 2, 2, 2, 66, 68, 5, 4, 3, 2, 67, 66,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 3, 3, 2, 2, 2, 69, 71, 5, 6, 4, 2,
	70, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 5, 3, 2, 2, 2, 74, 81, 5, 48, 25, 2, 75, 81, 5, 42, 22, 2,
	76, 81, 5, 54, 28, 2, 77, 81, 5, 36, 19, 2, 78, 81, 5, 60, 31, 2, 79, 81,
	5, 8, 5, 2, 80, 74, 3, 2, 2, 2, 80, 75, 3, 2, 2, 2, 80, 76, 3, 2, 2, 2,
	80, 77, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 7, 3, 2,
	2, 2, 82, 83, 7, 3, 2, 2, 83, 84, 7, 29, 2, 2, 84, 9, 3, 2, 2, 2, 85, 87,
	5, 12, 7, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2,
	88, 89, 3, 2, 2, 2, 89, 11, 3, 2, 2, 2, 90, 96, 5, 42, 22, 2, 91, 96, 5,
	16, 9, 2, 92, 96, 5, 14, 8, 2, 93, 96, 5, 22, 12, 2, 94, 96, 5, 18, 10,
	2, 95, 90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93,
	3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 13, 3, 2, 2, 2, 97, 98, 7, 4, 2, 2,
	98, 99, 5, 62, 32, 2, 99, 15, 3, 2, 2, 2, 100, 101, 7, 5, 2, 2, 101, 103,
	7, 6, 2, 2, 102, 104, 5, 10, 6, 2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 7, 7, 2, 2, 106, 17, 3, 2, 2, 2,
	107, 110, 5, 20, 11, 2, 108, 110, 5, 26, 14, 2, 109, 107, 3, 2, 2, 2, 109,
	108, 3, 2, 2, 2, 110, 19, 3, 2, 2, 2, 111, 112, 7, 8, 2, 2, 112, 113, 5,
	62, 32, 2, 113, 114, 7, 6, 2, 2, 114, 115, 5, 62, 32, 2, 115, 116, 7, 9,
	2, 2, 116, 124, 5, 12, 7, 2, 117, 118, 7, 10, 2, 2, 118, 119, 5, 62, 32,
	2, 119, 120, 7, 9, 2, 2, 120, 121, 5, 12, 7, 2, 121, 123, 3, 2, 2, 2, 122,
	117, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125,
	3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 129, 7, 10,
	2, 2, 128, 127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2,
	130, 131, 7, 7, 2, 2, 131, 21, 3, 2, 2, 2, 132, 133, 7, 29, 2, 2, 133,
	134, 7, 11, 2, 2, 134, 135, 5, 62, 32, 2, 135, 23, 3, 2, 2, 2, 136, 141,
	5, 62, 32, 2, 137, 138, 7, 10, 2, 2, 138, 140, 5, 62, 32, 2, 139, 137,
	3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 25, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 145, 7, 29, 2, 2,
	145, 147, 7, 12, 2, 2, 146, 148, 5, 24, 13, 2, 147, 146, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 7, 13, 2, 2, 150, 27,
	3, 2, 2, 2, 151, 152, 7, 29, 2, 2, 152, 29, 3, 2, 2, 2, 153, 158, 5, 28,
	15, 2, 154, 155, 7, 10, 2, 2, 155, 157, 5, 28, 15, 2, 156, 154, 3, 2, 2,
	2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159,
	31, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 163, 5, 34, 18, 2, 162, 161,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2,
	2, 2, 165, 33, 3, 2, 2, 2, 166, 168, 5, 38, 20, 2, 167, 166, 3, 2, 2, 2,
	167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 7, 29, 2, 2, 170,
	172, 7, 12, 2, 2, 171, 173, 5, 44, 23, 2, 172, 171, 3, 2, 2, 2, 172, 173,
	3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 177, 7, 13, 2, 2, 175, 176, 7, 14,
	2, 2, 176, 178, 5, 28, 15, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2,
	2, 178, 179, 3, 2, 2, 2, 179, 181, 7, 6, 2, 2, 180, 182, 5, 10, 6, 2, 181,
	180, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184,
	7, 7, 2, 2, 184, 35, 3, 2, 2, 2, 185, 186, 7, 15, 2, 2, 186, 189, 7, 29,
	2, 2, 187, 188, 7, 16, 2, 2, 188, 190, 5, 30, 16, 2, 189, 187, 3, 2, 2,
	2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193, 7, 6, 2, 2, 192,
	194, 5, 32, 17, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195,
	3, 2, 2, 2, 195, 196, 7, 7, 2, 2, 196, 37, 3, 2, 2, 2, 197, 198, 7, 17,
	2, 2, 198, 39, 3, 2, 2, 2, 199, 201, 5, 38, 20, 2, 200, 199, 3, 2, 2, 2,
	200, 201, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 205, 7, 29, 2, 2, 203,
	204, 7, 16, 2, 2, 204, 206, 5, 28, 15, 2, 205, 203, 3, 2, 2, 2, 205, 206,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 7, 11, 2, 2, 208, 209, 5, 62,
	32, 2, 209, 41, 3, 2, 2, 2, 210, 212, 7, 18, 2, 2, 211, 213, 7, 19, 2,
	2, 212, 211, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	219, 5, 40, 21, 2, 215, 216, 7, 10, 2, 2, 216, 218, 5, 40, 21, 2, 217,
	215, 3, 2, 2, 2, 218, 221, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220,
	3, 2, 2, 2, 220, 43, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 222, 227, 5, 46,
	24, 2, 223, 224, 7, 10, 2, 2, 224, 226, 5, 46, 24, 2, 225, 223, 3, 2, 2,
	2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	45, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 7, 29, 2, 2, 231, 232,
	7, 16, 2, 2, 232, 233, 5, 28, 15, 2, 233, 47, 3, 2, 2, 2, 234, 236, 7,
	20, 2, 2, 235, 237, 5, 38, 20, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2,
	2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 7, 29, 2, 2, 239, 241, 7, 12, 2,
	2, 240, 242, 5, 44, 23, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2,
	242, 243, 3, 2, 2, 2, 243, 246, 7, 13, 2, 2, 244, 245, 7, 14, 2, 2, 245,
	247, 5, 28, 15, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248,
	3, 2, 2, 2, 248, 250, 7, 6, 2, 2, 249, 251, 5, 10, 6, 2, 250, 249, 3, 2,
	2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 7, 2, 2,
	253, 49, 3, 2, 2, 2, 254, 256, 5, 52, 27, 2, 255, 254, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 51, 3,
	2, 2, 2, 259, 260, 7, 29, 2, 2, 260, 261, 7, 16, 2, 2, 261, 262, 5, 28,
	15, 2, 262, 53, 3, 2, 2, 2, 263, 265, 7, 21, 2, 2, 264, 266, 5, 38, 20,
	2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267,
	268, 7, 29, 2, 2, 268, 269, 7, 12, 2, 2, 269, 270, 5, 50, 26, 2, 270, 271,
	7, 13, 2, 2, 271, 55, 3, 2, 2, 2, 272, 274, 5, 58, 30, 2, 273, 272, 3,
	2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2,
	2, 276, 57, 3, 2, 2, 2, 277, 279, 5, 38, 20, 2, 278, 277, 3, 2, 2, 2, 278,
	279, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 281, 7, 29, 2, 2, 281, 283,
	7, 12, 2, 2, 282, 284, 5, 30, 16, 2, 283, 282, 3, 2, 2, 2, 283, 284, 3,
	2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 288, 7, 13, 2, 2, 286, 287, 7, 14,
	2, 2, 287, 289, 5, 28, 15, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2,
	2, 289, 59, 3, 2, 2, 2, 290, 291, 7, 22, 2, 2, 291, 292, 5, 38, 20, 2,
	292, 293, 7, 29, 2, 2, 293, 295, 7, 6, 2, 2, 294, 296, 5, 56, 29, 2, 295,
	294, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298,
	7, 7, 2, 2, 298, 61, 3, 2, 2, 2, 299, 300, 8, 32, 1, 2, 300, 301, 7, 12,
	2, 2, 301, 302, 5, 62, 32, 2, 302, 303, 7, 13, 2, 2, 303, 306, 3, 2, 2,
	2, 304, 306, 5, 64, 33, 2, 305, 299, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2,
	306, 318, 3, 2, 2, 2, 307, 308, 12, 6, 2, 2, 308, 309, 7, 23, 2, 2, 309,
	317, 5, 62, 32, 7, 310, 311, 12, 5, 2, 2, 311, 312, 9, 2, 2, 2, 312, 317,
	5, 62, 32, 6, 313, 314, 12, 4, 2, 2, 314, 315, 9, 3, 2, 2, 315, 317, 5,
	62, 32, 5, 316, 307, 3, 2, 2, 2, 316, 310, 3, 2, 2, 2, 316, 313, 3, 2,
	2, 2, 317, 320, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2,
	319, 63, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 321, 326, 5, 18, 10, 2, 322,
	326, 7, 30, 2, 2, 323, 326, 7, 29, 2, 2, 324, 326, 7, 31, 2, 2, 325, 321,
	3, 2, 2, 2, 325, 322, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 324, 3, 2,
	2, 2, 326, 65, 3, 2, 2, 2, 41, 67, 72, 80, 88, 95, 103, 109, 124, 128,
	141, 147, 158, 164, 167, 172, 177, 181, 189, 193, 200, 205, 212, 219, 227,
	236, 241, 246, 250, 257, 265, 275, 278, 283, 288, 295, 305, 316, 318, 325,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'import'", "'return'", "'loop'", "'{'", "'}'", "'match'", "'=>'",
	"','", "'='", "'('", "')'", "'->'", "'impl'", "':'", "'+'", "'let'", "'mut'",
	"'fn'", "'type'", "'trait'", "'^'", "'*'", "'/'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "WS", "COMMENT", "ID", "NUM", "STRING",
}

var ruleNames = []string{
	"prog", "topStatList", "topStat", "importStat", "statList", "stat", "returnStat",
	"loopStat", "exprStat", "matchRule", "assign", "exprList", "fnCall", "typePass",
	"typeList", "methodList", "method", "implBlock", "exportor", "define",
	"varDefine", "paramList", "param", "fnDefine", "attrList", "attr", "typeDefine",
	"tmethodList", "tmethod", "traitDefine", "expr", "factor",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ElzParser struct {
	*antlr.BaseParser
}

func NewElzParser(input antlr.TokenStream) *ElzParser {
	this := new(ElzParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Elz.g4"

	return this
}

// ElzParser tokens.
const (
	ElzParserEOF     = antlr.TokenEOF
	ElzParserT__0    = 1
	ElzParserT__1    = 2
	ElzParserT__2    = 3
	ElzParserT__3    = 4
	ElzParserT__4    = 5
	ElzParserT__5    = 6
	ElzParserT__6    = 7
	ElzParserT__7    = 8
	ElzParserT__8    = 9
	ElzParserT__9    = 10
	ElzParserT__10   = 11
	ElzParserT__11   = 12
	ElzParserT__12   = 13
	ElzParserT__13   = 14
	ElzParserT__14   = 15
	ElzParserT__15   = 16
	ElzParserT__16   = 17
	ElzParserT__17   = 18
	ElzParserT__18   = 19
	ElzParserT__19   = 20
	ElzParserT__20   = 21
	ElzParserT__21   = 22
	ElzParserT__22   = 23
	ElzParserT__23   = 24
	ElzParserWS      = 25
	ElzParserCOMMENT = 26
	ElzParserID      = 27
	ElzParserNUM     = 28
	ElzParserSTRING  = 29
)

// ElzParser rules.
const (
	ElzParserRULE_prog        = 0
	ElzParserRULE_topStatList = 1
	ElzParserRULE_topStat     = 2
	ElzParserRULE_importStat  = 3
	ElzParserRULE_statList    = 4
	ElzParserRULE_stat        = 5
	ElzParserRULE_returnStat  = 6
	ElzParserRULE_loopStat    = 7
	ElzParserRULE_exprStat    = 8
	ElzParserRULE_matchRule   = 9
	ElzParserRULE_assign      = 10
	ElzParserRULE_exprList    = 11
	ElzParserRULE_fnCall      = 12
	ElzParserRULE_typePass    = 13
	ElzParserRULE_typeList    = 14
	ElzParserRULE_methodList  = 15
	ElzParserRULE_method      = 16
	ElzParserRULE_implBlock   = 17
	ElzParserRULE_exportor    = 18
	ElzParserRULE_define      = 19
	ElzParserRULE_varDefine   = 20
	ElzParserRULE_paramList   = 21
	ElzParserRULE_param       = 22
	ElzParserRULE_fnDefine    = 23
	ElzParserRULE_attrList    = 24
	ElzParserRULE_attr        = 25
	ElzParserRULE_typeDefine  = 26
	ElzParserRULE_tmethodList = 27
	ElzParserRULE_tmethod     = 28
	ElzParserRULE_traitDefine = 29
	ElzParserRULE_expr        = 30
	ElzParserRULE_factor      = 31
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) TopStatList() ITopStatListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopStatListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITopStatListContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ElzParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ElzParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__0)|(1<<ElzParserT__12)|(1<<ElzParserT__15)|(1<<ElzParserT__17)|(1<<ElzParserT__18)|(1<<ElzParserT__19))) != 0 {
		{
			p.SetState(64)
			p.TopStatList()
		}

	}

	return localctx
}

// ITopStatListContext is an interface to support dynamic dispatch.
type ITopStatListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopStatListContext differentiates from other interfaces.
	IsTopStatListContext()
}

type TopStatListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopStatListContext() *TopStatListContext {
	var p = new(TopStatListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_topStatList
	return p
}

func (*TopStatListContext) IsTopStatListContext() {}

func NewTopStatListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopStatListContext {
	var p = new(TopStatListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_topStatList

	return p
}

func (s *TopStatListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopStatListContext) AllTopStat() []ITopStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopStatContext)(nil)).Elem())
	var tst = make([]ITopStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopStatContext)
		}
	}

	return tst
}

func (s *TopStatListContext) TopStat(i int) ITopStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopStatContext)
}

func (s *TopStatListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopStatListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopStatListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTopStatList(s)
	}
}

func (s *TopStatListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTopStatList(s)
	}
}

func (p *ElzParser) TopStatList() (localctx ITopStatListContext) {
	localctx = NewTopStatListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ElzParserRULE_topStatList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__0)|(1<<ElzParserT__12)|(1<<ElzParserT__15)|(1<<ElzParserT__17)|(1<<ElzParserT__18)|(1<<ElzParserT__19))) != 0) {
		{
			p.SetState(67)
			p.TopStat()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopStatContext is an interface to support dynamic dispatch.
type ITopStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopStatContext differentiates from other interfaces.
	IsTopStatContext()
}

type TopStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopStatContext() *TopStatContext {
	var p = new(TopStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_topStat
	return p
}

func (*TopStatContext) IsTopStatContext() {}

func NewTopStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopStatContext {
	var p = new(TopStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_topStat

	return p
}

func (s *TopStatContext) GetParser() antlr.Parser { return s.parser }

func (s *TopStatContext) FnDefine() IFnDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnDefineContext)
}

func (s *TopStatContext) VarDefine() IVarDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDefineContext)
}

func (s *TopStatContext) TypeDefine() ITypeDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefineContext)
}

func (s *TopStatContext) ImplBlock() IImplBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplBlockContext)
}

func (s *TopStatContext) TraitDefine() ITraitDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITraitDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITraitDefineContext)
}

func (s *TopStatContext) ImportStat() IImportStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportStatContext)
}

func (s *TopStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTopStat(s)
	}
}

func (s *TopStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTopStat(s)
	}
}

func (p *ElzParser) TopStat() (localctx ITopStatContext) {
	localctx = NewTopStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ElzParserRULE_topStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.FnDefine()
		}

	case ElzParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.VarDefine()
		}

	case ElzParserT__18:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.TypeDefine()
		}

	case ElzParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.ImplBlock()
		}

	case ElzParserT__19:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.TraitDefine()
		}

	case ElzParserT__0:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.ImportStat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportStatContext is an interface to support dynamic dispatch.
type IImportStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatContext differentiates from other interfaces.
	IsImportStatContext()
}

type ImportStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatContext() *ImportStatContext {
	var p = new(ImportStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_importStat
	return p
}

func (*ImportStatContext) IsImportStatContext() {}

func NewImportStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatContext {
	var p = new(ImportStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_importStat

	return p
}

func (s *ImportStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *ImportStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterImportStat(s)
	}
}

func (s *ImportStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitImportStat(s)
	}
}

func (p *ElzParser) ImportStat() (localctx IImportStatContext) {
	localctx = NewImportStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ElzParserRULE_importStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(ElzParserT__0)
	}
	{
		p.SetState(81)
		p.Match(ElzParserID)
	}

	return localctx
}

// IStatListContext is an interface to support dynamic dispatch.
type IStatListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatListContext differentiates from other interfaces.
	IsStatListContext()
}

type StatListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatListContext() *StatListContext {
	var p = new(StatListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_statList
	return p
}

func (*StatListContext) IsStatListContext() {}

func NewStatListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatListContext {
	var p = new(StatListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_statList

	return p
}

func (s *StatListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatListContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *StatListContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *StatListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterStatList(s)
	}
}

func (s *StatListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitStatList(s)
	}
}

func (p *ElzParser) StatList() (localctx IStatListContext) {
	localctx = NewStatListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ElzParserRULE_statList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__2)|(1<<ElzParserT__5)|(1<<ElzParserT__15)|(1<<ElzParserID))) != 0) {
		{
			p.SetState(83)
			p.Stat()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) VarDefine() IVarDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDefineContext)
}

func (s *StatContext) LoopStat() ILoopStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopStatContext)
}

func (s *StatContext) ReturnStat() IReturnStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatContext)
}

func (s *StatContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatContext) ExprStat() IExprStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *ElzParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ElzParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.VarDefine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.LoopStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.ReturnStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.ExprStat()
		}

	}

	return localctx
}

// IReturnStatContext is an interface to support dynamic dispatch.
type IReturnStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatContext differentiates from other interfaces.
	IsReturnStatContext()
}

type ReturnStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatContext() *ReturnStatContext {
	var p = new(ReturnStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_returnStat
	return p
}

func (*ReturnStatContext) IsReturnStatContext() {}

func NewReturnStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatContext {
	var p = new(ReturnStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_returnStat

	return p
}

func (s *ReturnStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterReturnStat(s)
	}
}

func (s *ReturnStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitReturnStat(s)
	}
}

func (p *ElzParser) ReturnStat() (localctx IReturnStatContext) {
	localctx = NewReturnStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ElzParserRULE_returnStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(ElzParserT__1)
	}
	{
		p.SetState(96)
		p.expr(0)
	}

	return localctx
}

// ILoopStatContext is an interface to support dynamic dispatch.
type ILoopStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopStatContext differentiates from other interfaces.
	IsLoopStatContext()
}

type LoopStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatContext() *LoopStatContext {
	var p = new(LoopStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_loopStat
	return p
}

func (*LoopStatContext) IsLoopStatContext() {}

func NewLoopStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatContext {
	var p = new(LoopStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_loopStat

	return p
}

func (s *LoopStatContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatContext) StatList() IStatListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatListContext)
}

func (s *LoopStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterLoopStat(s)
	}
}

func (s *LoopStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitLoopStat(s)
	}
}

func (p *ElzParser) LoopStat() (localctx ILoopStatContext) {
	localctx = NewLoopStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ElzParserRULE_loopStat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(ElzParserT__2)
	}
	{
		p.SetState(99)
		p.Match(ElzParserT__3)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__2)|(1<<ElzParserT__5)|(1<<ElzParserT__15)|(1<<ElzParserID))) != 0 {
		{
			p.SetState(100)
			p.StatList()
		}

	}
	{
		p.SetState(103)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IExprStatContext is an interface to support dynamic dispatch.
type IExprStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprStatContext differentiates from other interfaces.
	IsExprStatContext()
}

type ExprStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStatContext() *ExprStatContext {
	var p = new(ExprStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_exprStat
	return p
}

func (*ExprStatContext) IsExprStatContext() {}

func NewExprStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStatContext {
	var p = new(ExprStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_exprStat

	return p
}

func (s *ExprStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStatContext) MatchRule() IMatchRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchRuleContext)
}

func (s *ExprStatContext) FnCall() IFnCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnCallContext)
}

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterExprStat(s)
	}
}

func (s *ExprStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitExprStat(s)
	}
}

func (p *ElzParser) ExprStat() (localctx IExprStatContext) {
	localctx = NewExprStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ElzParserRULE_exprStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.MatchRule()
		}

	case ElzParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.FnCall()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMatchRuleContext is an interface to support dynamic dispatch.
type IMatchRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchRuleContext differentiates from other interfaces.
	IsMatchRuleContext()
}

type MatchRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchRuleContext() *MatchRuleContext {
	var p = new(MatchRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_matchRule
	return p
}

func (*MatchRuleContext) IsMatchRuleContext() {}

func NewMatchRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchRuleContext {
	var p = new(MatchRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_matchRule

	return p
}

func (s *MatchRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchRuleContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MatchRuleContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchRuleContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *MatchRuleContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *MatchRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterMatchRule(s)
	}
}

func (s *MatchRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitMatchRule(s)
	}
}

func (p *ElzParser) MatchRule() (localctx IMatchRuleContext) {
	localctx = NewMatchRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ElzParserRULE_matchRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(ElzParserT__5)
	}
	{
		p.SetState(110)
		p.expr(0)
	}
	{
		p.SetState(111)
		p.Match(ElzParserT__3)
	}
	{
		p.SetState(112)
		p.expr(0)
	}
	{
		p.SetState(113)
		p.Match(ElzParserT__6)
	}
	{
		p.SetState(114)
		p.Stat()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(115)
				p.Match(ElzParserT__7)
			}
			{
				p.SetState(116)
				p.expr(0)
			}
			{
				p.SetState(117)
				p.Match(ElzParserT__6)
			}
			{
				p.SetState(118)
				p.Stat()
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__7 {
		{
			p.SetState(125)
			p.Match(ElzParserT__7)
		}

	}
	{
		p.SetState(128)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *ElzParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ElzParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(ElzParserID)
	}
	{
		p.SetState(131)
		p.Match(ElzParserT__8)
	}
	{
		p.SetState(132)
		p.expr(0)
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *ElzParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ElzParserRULE_exprList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.expr(0)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__7 {
		{
			p.SetState(135)
			p.Match(ElzParserT__7)
		}
		{
			p.SetState(136)
			p.expr(0)
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFnCallContext is an interface to support dynamic dispatch.
type IFnCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnCallContext differentiates from other interfaces.
	IsFnCallContext()
}

type FnCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnCallContext() *FnCallContext {
	var p = new(FnCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_fnCall
	return p
}

func (*FnCallContext) IsFnCallContext() {}

func NewFnCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnCallContext {
	var p = new(FnCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_fnCall

	return p
}

func (s *FnCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FnCallContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *FnCallContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *FnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterFnCall(s)
	}
}

func (s *FnCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitFnCall(s)
	}
}

func (p *ElzParser) FnCall() (localctx IFnCallContext) {
	localctx = NewFnCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ElzParserRULE_fnCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(ElzParserID)
	}
	{
		p.SetState(143)
		p.Match(ElzParserT__9)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__5)|(1<<ElzParserT__9)|(1<<ElzParserID)|(1<<ElzParserNUM)|(1<<ElzParserSTRING))) != 0 {
		{
			p.SetState(144)
			p.ExprList()
		}

	}
	{
		p.SetState(147)
		p.Match(ElzParserT__10)
	}

	return localctx
}

// ITypePassContext is an interface to support dynamic dispatch.
type ITypePassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypePassContext differentiates from other interfaces.
	IsTypePassContext()
}

type TypePassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypePassContext() *TypePassContext {
	var p = new(TypePassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_typePass
	return p
}

func (*TypePassContext) IsTypePassContext() {}

func NewTypePassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypePassContext {
	var p = new(TypePassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_typePass

	return p
}

func (s *TypePassContext) GetParser() antlr.Parser { return s.parser }

func (s *TypePassContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TypePassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypePassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypePassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTypePass(s)
	}
}

func (s *TypePassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTypePass(s)
	}
}

func (p *ElzParser) TypePass() (localctx ITypePassContext) {
	localctx = NewTypePassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ElzParserRULE_typePass)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(ElzParserID)
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypePass() []ITypePassContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypePassContext)(nil)).Elem())
	var tst = make([]ITypePassContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypePassContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypePass(i int) ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (p *ElzParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ElzParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.TypePass()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__7 {
		{
			p.SetState(152)
			p.Match(ElzParserT__7)
		}
		{
			p.SetState(153)
			p.TypePass()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodListContext is an interface to support dynamic dispatch.
type IMethodListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodListContext differentiates from other interfaces.
	IsMethodListContext()
}

type MethodListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodListContext() *MethodListContext {
	var p = new(MethodListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_methodList
	return p
}

func (*MethodListContext) IsMethodListContext() {}

func NewMethodListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodListContext {
	var p = new(MethodListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_methodList

	return p
}

func (s *MethodListContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodListContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *MethodListContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *MethodListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterMethodList(s)
	}
}

func (s *MethodListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitMethodList(s)
	}
}

func (p *ElzParser) MethodList() (localctx IMethodListContext) {
	localctx = NewMethodListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ElzParserRULE_methodList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__14 || _la == ElzParserID {
		{
			p.SetState(159)
			p.Method()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *MethodContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *MethodContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *MethodContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *MethodContext) StatList() IStatListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatListContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *ElzParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ElzParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 {
		{
			p.SetState(164)
			p.Exportor()
		}

	}
	{
		p.SetState(167)
		p.Match(ElzParserID)
	}
	{
		p.SetState(168)
		p.Match(ElzParserT__9)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(169)
			p.ParamList()
		}

	}
	{
		p.SetState(172)
		p.Match(ElzParserT__10)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__11 {
		{
			p.SetState(173)
			p.Match(ElzParserT__11)
		}
		{
			p.SetState(174)
			p.TypePass()
		}

	}
	{
		p.SetState(177)
		p.Match(ElzParserT__3)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__2)|(1<<ElzParserT__5)|(1<<ElzParserT__15)|(1<<ElzParserID))) != 0 {
		{
			p.SetState(178)
			p.StatList()
		}

	}
	{
		p.SetState(181)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IImplBlockContext is an interface to support dynamic dispatch.
type IImplBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplBlockContext differentiates from other interfaces.
	IsImplBlockContext()
}

type ImplBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplBlockContext() *ImplBlockContext {
	var p = new(ImplBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_implBlock
	return p
}

func (*ImplBlockContext) IsImplBlockContext() {}

func NewImplBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplBlockContext {
	var p = new(ImplBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_implBlock

	return p
}

func (s *ImplBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *ImplBlockContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *ImplBlockContext) MethodList() IMethodListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodListContext)
}

func (s *ImplBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterImplBlock(s)
	}
}

func (s *ImplBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitImplBlock(s)
	}
}

func (p *ElzParser) ImplBlock() (localctx IImplBlockContext) {
	localctx = NewImplBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ElzParserRULE_implBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(ElzParserT__12)
	}
	{
		p.SetState(184)
		p.Match(ElzParserID)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__13 {
		{
			p.SetState(185)
			p.Match(ElzParserT__13)
		}
		{
			p.SetState(186)
			p.TypeList()
		}

	}
	{
		p.SetState(189)
		p.Match(ElzParserT__3)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 || _la == ElzParserID {
		{
			p.SetState(190)
			p.MethodList()
		}

	}
	{
		p.SetState(193)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IExportorContext is an interface to support dynamic dispatch.
type IExportorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportorContext differentiates from other interfaces.
	IsExportorContext()
}

type ExportorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportorContext() *ExportorContext {
	var p = new(ExportorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_exportor
	return p
}

func (*ExportorContext) IsExportorContext() {}

func NewExportorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportorContext {
	var p = new(ExportorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_exportor

	return p
}

func (s *ExportorContext) GetParser() antlr.Parser { return s.parser }
func (s *ExportorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterExportor(s)
	}
}

func (s *ExportorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitExportor(s)
	}
}

func (p *ElzParser) Exportor() (localctx IExportorContext) {
	localctx = NewExportorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ElzParserRULE_exportor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(ElzParserT__14)
	}

	return localctx
}

// IDefineContext is an interface to support dynamic dispatch.
type IDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineContext differentiates from other interfaces.
	IsDefineContext()
}

type DefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineContext() *DefineContext {
	var p = new(DefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_define
	return p
}

func (*DefineContext) IsDefineContext() {}

func NewDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineContext {
	var p = new(DefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_define

	return p
}

func (s *DefineContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *DefineContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *DefineContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *DefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterDefine(s)
	}
}

func (s *DefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitDefine(s)
	}
}

func (p *ElzParser) Define() (localctx IDefineContext) {
	localctx = NewDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ElzParserRULE_define)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 {
		{
			p.SetState(197)
			p.Exportor()
		}

	}
	{
		p.SetState(200)
		p.Match(ElzParserID)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__13 {
		{
			p.SetState(201)
			p.Match(ElzParserT__13)
		}
		{
			p.SetState(202)
			p.TypePass()
		}

	}
	{
		p.SetState(205)
		p.Match(ElzParserT__8)
	}
	{
		p.SetState(206)
		p.expr(0)
	}

	return localctx
}

// IVarDefineContext is an interface to support dynamic dispatch.
type IVarDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMut returns the mut token.
	GetMut() antlr.Token

	// SetMut sets the mut token.
	SetMut(antlr.Token)

	// IsVarDefineContext differentiates from other interfaces.
	IsVarDefineContext()
}

type VarDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mut    antlr.Token
}

func NewEmptyVarDefineContext() *VarDefineContext {
	var p = new(VarDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_varDefine
	return p
}

func (*VarDefineContext) IsVarDefineContext() {}

func NewVarDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefineContext {
	var p = new(VarDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_varDefine

	return p
}

func (s *VarDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefineContext) GetMut() antlr.Token { return s.mut }

func (s *VarDefineContext) SetMut(v antlr.Token) { s.mut = v }

func (s *VarDefineContext) AllDefine() []IDefineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefineContext)(nil)).Elem())
	var tst = make([]IDefineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefineContext)
		}
	}

	return tst
}

func (s *VarDefineContext) Define(i int) IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *VarDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterVarDefine(s)
	}
}

func (s *VarDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitVarDefine(s)
	}
}

func (p *ElzParser) VarDefine() (localctx IVarDefineContext) {
	localctx = NewVarDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ElzParserRULE_varDefine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(ElzParserT__15)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(209)

			var _m = p.Match(ElzParserT__16)

			localctx.(*VarDefineContext).mut = _m
		}

	}
	{
		p.SetState(212)
		p.Define()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(213)
				p.Match(ElzParserT__7)
			}
			{
				p.SetState(214)
				p.Define()
			}

		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *ElzParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ElzParserRULE_paramList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Param()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__7 {
		{
			p.SetState(221)
			p.Match(ElzParserT__7)
		}
		{
			p.SetState(222)
			p.Param()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *ParamContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *ElzParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ElzParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(ElzParserID)
	}
	{
		p.SetState(229)
		p.Match(ElzParserT__13)
	}
	{
		p.SetState(230)
		p.TypePass()
	}

	return localctx
}

// IFnDefineContext is an interface to support dynamic dispatch.
type IFnDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnDefineContext differentiates from other interfaces.
	IsFnDefineContext()
}

type FnDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnDefineContext() *FnDefineContext {
	var p = new(FnDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_fnDefine
	return p
}

func (*FnDefineContext) IsFnDefineContext() {}

func NewFnDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnDefineContext {
	var p = new(FnDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_fnDefine

	return p
}

func (s *FnDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *FnDefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *FnDefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *FnDefineContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FnDefineContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *FnDefineContext) StatList() IStatListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatListContext)
}

func (s *FnDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterFnDefine(s)
	}
}

func (s *FnDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitFnDefine(s)
	}
}

func (p *ElzParser) FnDefine() (localctx IFnDefineContext) {
	localctx = NewFnDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ElzParserRULE_fnDefine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(ElzParserT__17)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 {
		{
			p.SetState(233)
			p.Exportor()
		}

	}
	{
		p.SetState(236)
		p.Match(ElzParserID)
	}
	{
		p.SetState(237)
		p.Match(ElzParserT__9)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(238)
			p.ParamList()
		}

	}
	{
		p.SetState(241)
		p.Match(ElzParserT__10)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__11 {
		{
			p.SetState(242)
			p.Match(ElzParserT__11)
		}
		{
			p.SetState(243)
			p.TypePass()
		}

	}
	{
		p.SetState(246)
		p.Match(ElzParserT__3)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__2)|(1<<ElzParserT__5)|(1<<ElzParserT__15)|(1<<ElzParserID))) != 0 {
		{
			p.SetState(247)
			p.StatList()
		}

	}
	{
		p.SetState(250)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IAttrListContext is an interface to support dynamic dispatch.
type IAttrListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrListContext differentiates from other interfaces.
	IsAttrListContext()
}

type AttrListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrListContext() *AttrListContext {
	var p = new(AttrListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_attrList
	return p
}

func (*AttrListContext) IsAttrListContext() {}

func NewAttrListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrListContext {
	var p = new(AttrListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_attrList

	return p
}

func (s *AttrListContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrListContext) AllAttr() []IAttrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttrContext)(nil)).Elem())
	var tst = make([]IAttrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttrContext)
		}
	}

	return tst
}

func (s *AttrListContext) Attr(i int) IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *AttrListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAttrList(s)
	}
}

func (s *AttrListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAttrList(s)
	}
}

func (p *ElzParser) AttrList() (localctx IAttrListContext) {
	localctx = NewAttrListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ElzParserRULE_attrList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserID {
		{
			p.SetState(252)
			p.Attr()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *AttrContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *ElzParser) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ElzParserRULE_attr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(ElzParserID)
	}
	{
		p.SetState(258)
		p.Match(ElzParserT__13)
	}
	{
		p.SetState(259)
		p.TypePass()
	}

	return localctx
}

// ITypeDefineContext is an interface to support dynamic dispatch.
type ITypeDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefineContext differentiates from other interfaces.
	IsTypeDefineContext()
}

type TypeDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefineContext() *TypeDefineContext {
	var p = new(TypeDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_typeDefine
	return p
}

func (*TypeDefineContext) IsTypeDefineContext() {}

func NewTypeDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefineContext {
	var p = new(TypeDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_typeDefine

	return p
}

func (s *TypeDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TypeDefineContext) AttrList() IAttrListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrListContext)
}

func (s *TypeDefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *TypeDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTypeDefine(s)
	}
}

func (s *TypeDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTypeDefine(s)
	}
}

func (p *ElzParser) TypeDefine() (localctx ITypeDefineContext) {
	localctx = NewTypeDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ElzParserRULE_typeDefine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(ElzParserT__18)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 {
		{
			p.SetState(262)
			p.Exportor()
		}

	}
	{
		p.SetState(265)
		p.Match(ElzParserID)
	}
	{
		p.SetState(266)
		p.Match(ElzParserT__9)
	}
	{
		p.SetState(267)
		p.AttrList()
	}
	{
		p.SetState(268)
		p.Match(ElzParserT__10)
	}

	return localctx
}

// ITmethodListContext is an interface to support dynamic dispatch.
type ITmethodListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmethodListContext differentiates from other interfaces.
	IsTmethodListContext()
}

type TmethodListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmethodListContext() *TmethodListContext {
	var p = new(TmethodListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_tmethodList
	return p
}

func (*TmethodListContext) IsTmethodListContext() {}

func NewTmethodListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmethodListContext {
	var p = new(TmethodListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_tmethodList

	return p
}

func (s *TmethodListContext) GetParser() antlr.Parser { return s.parser }

func (s *TmethodListContext) AllTmethod() []ITmethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITmethodContext)(nil)).Elem())
	var tst = make([]ITmethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITmethodContext)
		}
	}

	return tst
}

func (s *TmethodListContext) Tmethod(i int) ITmethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITmethodContext)
}

func (s *TmethodListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmethodListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmethodListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTmethodList(s)
	}
}

func (s *TmethodListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTmethodList(s)
	}
}

func (p *ElzParser) TmethodList() (localctx ITmethodListContext) {
	localctx = NewTmethodListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ElzParserRULE_tmethodList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__14 || _la == ElzParserID {
		{
			p.SetState(270)
			p.Tmethod()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmethodContext is an interface to support dynamic dispatch.
type ITmethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmethodContext differentiates from other interfaces.
	IsTmethodContext()
}

type TmethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmethodContext() *TmethodContext {
	var p = new(TmethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_tmethod
	return p
}

func (*TmethodContext) IsTmethodContext() {}

func NewTmethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmethodContext {
	var p = new(TmethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_tmethod

	return p
}

func (s *TmethodContext) GetParser() antlr.Parser { return s.parser }

func (s *TmethodContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TmethodContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *TmethodContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *TmethodContext) TypePass() ITypePassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypePassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypePassContext)
}

func (s *TmethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTmethod(s)
	}
}

func (s *TmethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTmethod(s)
	}
}

func (p *ElzParser) Tmethod() (localctx ITmethodContext) {
	localctx = NewTmethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ElzParserRULE_tmethod)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 {
		{
			p.SetState(275)
			p.Exportor()
		}

	}
	{
		p.SetState(278)
		p.Match(ElzParserID)
	}
	{
		p.SetState(279)
		p.Match(ElzParserT__9)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(280)
			p.TypeList()
		}

	}
	{
		p.SetState(283)
		p.Match(ElzParserT__10)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__11 {
		{
			p.SetState(284)
			p.Match(ElzParserT__11)
		}
		{
			p.SetState(285)
			p.TypePass()
		}

	}

	return localctx
}

// ITraitDefineContext is an interface to support dynamic dispatch.
type ITraitDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTraitDefineContext differentiates from other interfaces.
	IsTraitDefineContext()
}

type TraitDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTraitDefineContext() *TraitDefineContext {
	var p = new(TraitDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_traitDefine
	return p
}

func (*TraitDefineContext) IsTraitDefineContext() {}

func NewTraitDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TraitDefineContext {
	var p = new(TraitDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_traitDefine

	return p
}

func (s *TraitDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *TraitDefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *TraitDefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TraitDefineContext) TmethodList() ITmethodListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmethodListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmethodListContext)
}

func (s *TraitDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TraitDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TraitDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTraitDefine(s)
	}
}

func (s *TraitDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTraitDefine(s)
	}
}

func (p *ElzParser) TraitDefine() (localctx ITraitDefineContext) {
	localctx = NewTraitDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ElzParserRULE_traitDefine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(ElzParserT__19)
	}
	{
		p.SetState(289)
		p.Exportor()
	}
	{
		p.SetState(290)
		p.Match(ElzParserID)
	}
	{
		p.SetState(291)
		p.Match(ElzParserT__3)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__14 || _la == ElzParserID {
		{
			p.SetState(292)
			p.TmethodList()
		}

	}
	{
		p.SetState(295)
		p.Match(ElzParserT__4)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ElzParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ElzParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, ElzParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__9:
		{
			p.SetState(298)
			p.Match(ElzParserT__9)
		}
		{
			p.SetState(299)
			p.expr(0)
		}
		{
			p.SetState(300)
			p.Match(ElzParserT__10)
		}

	case ElzParserT__5, ElzParserID, ElzParserNUM, ElzParserSTRING:
		{
			p.SetState(302)
			p.Factor()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(306)

					var _m = p.Match(ElzParserT__20)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(307)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(309)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ElzParserT__21 || _la == ElzParserT__22) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(310)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(312)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ElzParserT__14 || _la == ElzParserT__23) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(313)
					p.expr(3)
				}

			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) ExprStat() IExprStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprStatContext)
}

func (s *FactorContext) NUM() antlr.TerminalNode {
	return s.GetToken(ElzParserNUM, 0)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *FactorContext) STRING() antlr.TerminalNode {
	return s.GetToken(ElzParserSTRING, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *ElzParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ElzParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.ExprStat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(ElzParserNUM)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(321)
			p.Match(ElzParserID)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(322)
			p.Match(ElzParserSTRING)
		}

	}

	return localctx
}

func (p *ElzParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 30:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ElzParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
