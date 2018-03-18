// Code generated from Elz.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 368,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 3, 2, 5, 2, 72, 10, 2, 3, 3, 6, 3, 75, 10, 3, 13,
	3, 14, 3, 76, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 85, 10, 4, 3, 5,
	3, 5, 3, 5, 7, 5, 90, 10, 5, 12, 5, 14, 5, 93, 11, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 6, 7, 101, 10, 7, 13, 7, 14, 7, 102, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 110, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5,
	10, 118, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 124, 10, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7,
	12, 137, 10, 12, 12, 12, 14, 12, 140, 11, 12, 3, 12, 5, 12, 143, 10, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 154,
	10, 14, 12, 14, 14, 14, 157, 11, 14, 3, 15, 3, 15, 3, 15, 5, 15, 162, 10,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 171, 10, 17,
	12, 17, 14, 17, 174, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	5, 18, 182, 10, 18, 3, 19, 6, 19, 185, 10, 19, 13, 19, 14, 19, 186, 3,
	20, 5, 20, 190, 10, 20, 3, 20, 3, 20, 3, 20, 5, 20, 195, 10, 20, 3, 20,
	3, 20, 3, 20, 5, 20, 200, 10, 20, 3, 20, 3, 20, 5, 20, 204, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 212, 10, 21, 3, 21, 3, 21,
	5, 21, 216, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 5, 23, 223, 10,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 230, 10, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 5, 25, 237, 10, 25, 3, 25, 3, 25, 3, 25, 7, 25, 242,
	10, 25, 12, 25, 14, 25, 245, 11, 25, 3, 26, 3, 26, 3, 26, 7, 26, 250, 10,
	26, 12, 26, 14, 26, 253, 11, 26, 3, 27, 3, 27, 3, 27, 5, 27, 258, 10, 27,
	3, 28, 5, 28, 261, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 267, 10,
	28, 3, 28, 3, 28, 3, 28, 5, 28, 272, 10, 28, 3, 28, 3, 28, 5, 28, 276,
	10, 28, 3, 28, 3, 28, 3, 29, 6, 29, 281, 10, 29, 13, 29, 14, 29, 282, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 292, 10, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 6, 32, 300, 10, 32, 13, 32, 14, 32,
	301, 3, 33, 3, 33, 3, 33, 5, 33, 307, 10, 33, 3, 33, 3, 33, 3, 33, 5, 33,
	312, 10, 33, 3, 34, 5, 34, 315, 10, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5,
	34, 321, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 5, 35, 334, 10, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 7, 35, 363, 10, 35, 12, 35, 14, 35, 366, 11, 35, 3, 35, 2, 3,
	68, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 2,
	6, 3, 2, 26, 27, 4, 2, 19, 19, 28, 28, 3, 2, 29, 32, 3, 2, 35, 36, 2, 390,
	2, 71, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 84, 3, 2, 2, 2, 8, 86, 3, 2, 2,
	2, 10, 94, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 111,
	3, 2, 2, 2, 18, 114, 3, 2, 2, 2, 20, 123, 3, 2, 2, 2, 22, 125, 3, 2, 2,
	2, 24, 146, 3, 2, 2, 2, 26, 150, 3, 2, 2, 2, 28, 158, 3, 2, 2, 2, 30, 165,
	3, 2, 2, 2, 32, 167, 3, 2, 2, 2, 34, 175, 3, 2, 2, 2, 36, 184, 3, 2, 2,
	2, 38, 189, 3, 2, 2, 2, 40, 207, 3, 2, 2, 2, 42, 219, 3, 2, 2, 2, 44, 222,
	3, 2, 2, 2, 46, 226, 3, 2, 2, 2, 48, 234, 3, 2, 2, 2, 50, 246, 3, 2, 2,
	2, 52, 254, 3, 2, 2, 2, 54, 260, 3, 2, 2, 2, 56, 280, 3, 2, 2, 2, 58, 284,
	3, 2, 2, 2, 60, 289, 3, 2, 2, 2, 62, 299, 3, 2, 2, 2, 64, 303, 3, 2, 2,
	2, 66, 314, 3, 2, 2, 2, 68, 333, 3, 2, 2, 2, 70, 72, 5, 4, 3, 2, 71, 70,
	3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 3, 3, 2, 2, 2, 73, 75, 5, 6, 4, 2,
	74, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3,
	2, 2, 2, 77, 5, 3, 2, 2, 2, 78, 85, 5, 54, 28, 2, 79, 85, 5, 44, 23, 2,
	80, 85, 5, 60, 31, 2, 81, 85, 5, 40, 21, 2, 82, 85, 5, 66, 34, 2, 83, 85,
	5, 10, 6, 2, 84, 78, 3, 2, 2, 2, 84, 79, 3, 2, 2, 2, 84, 80, 3, 2, 2, 2,
	84, 81, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 7, 3, 2,
	2, 2, 86, 91, 7, 40, 2, 2, 87, 88, 7, 3, 2, 2, 88, 90, 7, 40, 2, 2, 89,
	87, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 9, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 7, 4, 2, 2, 95, 96, 7,
	5, 2, 2, 96, 97, 5, 8, 5, 2, 97, 98, 7, 6, 2, 2, 98, 11, 3, 2, 2, 2, 99,
	101, 5, 14, 8, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100,
	3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 13, 3, 2, 2, 2, 104, 110, 5, 48,
	25, 2, 105, 110, 5, 18, 10, 2, 106, 110, 5, 16, 9, 2, 107, 110, 5, 24,
	13, 2, 108, 110, 5, 20, 11, 2, 109, 104, 3, 2, 2, 2, 109, 105, 3, 2, 2,
	2, 109, 106, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110,
	15, 3, 2, 2, 2, 111, 112, 7, 7, 2, 2, 112, 113, 5, 68, 35, 2, 113, 17,
	3, 2, 2, 2, 114, 115, 7, 8, 2, 2, 115, 117, 7, 9, 2, 2, 116, 118, 5, 12,
	7, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2,
	119, 120, 7, 10, 2, 2, 120, 19, 3, 2, 2, 2, 121, 124, 5, 22, 12, 2, 122,
	124, 5, 28, 15, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 21,
	3, 2, 2, 2, 125, 126, 7, 11, 2, 2, 126, 127, 5, 68, 35, 2, 127, 128, 7,
	9, 2, 2, 128, 129, 5, 68, 35, 2, 129, 130, 7, 12, 2, 2, 130, 138, 5, 14,
	8, 2, 131, 132, 7, 13, 2, 2, 132, 133, 5, 68, 35, 2, 133, 134, 7, 12, 2,
	2, 134, 135, 5, 14, 8, 2, 135, 137, 3, 2, 2, 2, 136, 131, 3, 2, 2, 2, 137,
	140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 142,
	3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 143, 7, 13, 2, 2, 142, 141, 3, 2,
	2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 10, 2, 2,
	145, 23, 3, 2, 2, 2, 146, 147, 7, 40, 2, 2, 147, 148, 7, 14, 2, 2, 148,
	149, 5, 68, 35, 2, 149, 25, 3, 2, 2, 2, 150, 155, 5, 68, 35, 2, 151, 152,
	7, 13, 2, 2, 152, 154, 5, 68, 35, 2, 153, 151, 3, 2, 2, 2, 154, 157, 3,
	2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 27, 3, 2, 2,
	2, 157, 155, 3, 2, 2, 2, 158, 159, 7, 40, 2, 2, 159, 161, 7, 5, 2, 2, 160,
	162, 5, 26, 14, 2, 161, 160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 164, 7, 6, 2, 2, 164, 29, 3, 2, 2, 2, 165, 166, 7, 40,
	2, 2, 166, 31, 3, 2, 2, 2, 167, 172, 5, 30, 16, 2, 168, 169, 7, 13, 2,
	2, 169, 171, 5, 30, 16, 2, 170, 168, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 33, 3, 2, 2, 2, 174, 172,
	3, 2, 2, 2, 175, 176, 7, 15, 2, 2, 176, 181, 7, 40, 2, 2, 177, 178, 7,
	5, 2, 2, 178, 179, 5, 68, 35, 2, 179, 180, 7, 6, 2, 2, 180, 182, 3, 2,
	2, 2, 181, 177, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 35, 3, 2, 2, 2,
	183, 185, 5, 38, 20, 2, 184, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 37, 3, 2, 2, 2, 188, 190, 5,
	42, 22, 2, 189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2,
	2, 2, 191, 192, 7, 40, 2, 2, 192, 194, 7, 5, 2, 2, 193, 195, 5, 50, 26,
	2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	199, 7, 6, 2, 2, 197, 198, 7, 16, 2, 2, 198, 200, 5, 30, 16, 2, 199, 197,
	3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 7, 9,
	2, 2, 202, 204, 5, 12, 7, 2, 203, 202, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2,
	204, 205, 3, 2, 2, 2, 205, 206, 7, 10, 2, 2, 206, 39, 3, 2, 2, 2, 207,
	208, 7, 17, 2, 2, 208, 211, 7, 40, 2, 2, 209, 210, 7, 18, 2, 2, 210, 212,
	5, 32, 17, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3,
	2, 2, 2, 213, 215, 7, 9, 2, 2, 214, 216, 5, 36, 19, 2, 215, 214, 3, 2,
	2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 7, 10, 2, 2,
	218, 41, 3, 2, 2, 2, 219, 220, 7, 19, 2, 2, 220, 43, 3, 2, 2, 2, 221, 223,
	5, 42, 22, 2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 3,
	2, 2, 2, 224, 225, 5, 46, 24, 2, 225, 45, 3, 2, 2, 2, 226, 229, 7, 40,
	2, 2, 227, 228, 7, 18, 2, 2, 228, 230, 5, 30, 16, 2, 229, 227, 3, 2, 2,
	2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 7, 14, 2, 2, 232,
	233, 5, 68, 35, 2, 233, 47, 3, 2, 2, 2, 234, 236, 7, 20, 2, 2, 235, 237,
	7, 21, 2, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 3, 2,
	2, 2, 238, 243, 5, 46, 24, 2, 239, 240, 7, 13, 2, 2, 240, 242, 5, 46, 24,
	2, 241, 239, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243,
	244, 3, 2, 2, 2, 244, 49, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 251, 5,
	52, 27, 2, 247, 248, 7, 13, 2, 2, 248, 250, 5, 52, 27, 2, 249, 247, 3,
	2, 2, 2, 250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2,
	2, 252, 51, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 257, 7, 40, 2, 2, 255,
	256, 7, 18, 2, 2, 256, 258, 5, 30, 16, 2, 257, 255, 3, 2, 2, 2, 257, 258,
	3, 2, 2, 2, 258, 53, 3, 2, 2, 2, 259, 261, 5, 42, 22, 2, 260, 259, 3, 2,
	2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 263, 7, 22, 2, 2,
	263, 264, 7, 40, 2, 2, 264, 266, 7, 5, 2, 2, 265, 267, 5, 50, 26, 2, 266,
	265, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 271,
	7, 6, 2, 2, 269, 270, 7, 16, 2, 2, 270, 272, 5, 30, 16, 2, 271, 269, 3,
	2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 275, 7, 9, 2,
	2, 274, 276, 5, 12, 7, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276,
	277, 3, 2, 2, 2, 277, 278, 7, 10, 2, 2, 278, 55, 3, 2, 2, 2, 279, 281,
	5, 58, 30, 2, 280, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 280, 3,
	2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 57, 3, 2, 2, 2, 284, 285, 5, 42, 22,
	2, 285, 286, 7, 40, 2, 2, 286, 287, 7, 18, 2, 2, 287, 288, 5, 30, 16, 2,
	288, 59, 3, 2, 2, 2, 289, 291, 7, 23, 2, 2, 290, 292, 5, 42, 22, 2, 291,
	290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294,
	7, 40, 2, 2, 294, 295, 7, 5, 2, 2, 295, 296, 5, 56, 29, 2, 296, 297, 7,
	6, 2, 2, 297, 61, 3, 2, 2, 2, 298, 300, 5, 64, 33, 2, 299, 298, 3, 2, 2,
	2, 300, 301, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302,
	63, 3, 2, 2, 2, 303, 304, 7, 40, 2, 2, 304, 306, 7, 5, 2, 2, 305, 307,
	5, 32, 17, 2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3,
	2, 2, 2, 308, 311, 7, 6, 2, 2, 309, 310, 7, 16, 2, 2, 310, 312, 5, 30,
	16, 2, 311, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 65, 3, 2, 2, 2,
	313, 315, 5, 42, 22, 2, 314, 313, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315,
	316, 3, 2, 2, 2, 316, 317, 7, 24, 2, 2, 317, 318, 7, 40, 2, 2, 318, 320,
	7, 9, 2, 2, 319, 321, 5, 62, 32, 2, 320, 319, 3, 2, 2, 2, 320, 321, 3,
	2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 7, 10, 2, 2, 323, 67, 3, 2, 2,
	2, 324, 325, 8, 35, 1, 2, 325, 326, 7, 5, 2, 2, 326, 327, 5, 68, 35, 2,
	327, 328, 7, 6, 2, 2, 328, 334, 3, 2, 2, 2, 329, 334, 5, 20, 11, 2, 330,
	334, 7, 41, 2, 2, 331, 334, 7, 40, 2, 2, 332, 334, 7, 42, 2, 2, 333, 324,
	3, 2, 2, 2, 333, 329, 3, 2, 2, 2, 333, 330, 3, 2, 2, 2, 333, 331, 3, 2,
	2, 2, 333, 332, 3, 2, 2, 2, 334, 364, 3, 2, 2, 2, 335, 336, 12, 15, 2,
	2, 336, 337, 7, 25, 2, 2, 337, 363, 5, 68, 35, 16, 338, 339, 12, 14, 2,
	2, 339, 340, 9, 2, 2, 2, 340, 363, 5, 68, 35, 15, 341, 342, 12, 13, 2,
	2, 342, 343, 9, 3, 2, 2, 343, 363, 5, 68, 35, 14, 344, 345, 12, 12, 2,
	2, 345, 346, 9, 4, 2, 2, 346, 363, 5, 68, 35, 13, 347, 348, 12, 11, 2,
	2, 348, 349, 7, 33, 2, 2, 349, 363, 5, 68, 35, 12, 350, 351, 12, 10, 2,
	2, 351, 352, 7, 34, 2, 2, 352, 363, 5, 68, 35, 11, 353, 354, 12, 9, 2,
	2, 354, 355, 9, 5, 2, 2, 355, 363, 5, 68, 35, 10, 356, 357, 12, 8, 2, 2,
	357, 358, 7, 37, 2, 2, 358, 359, 5, 68, 35, 2, 359, 360, 7, 18, 2, 2, 360,
	361, 5, 68, 35, 9, 361, 363, 3, 2, 2, 2, 362, 335, 3, 2, 2, 2, 362, 338,
	3, 2, 2, 2, 362, 341, 3, 2, 2, 2, 362, 344, 3, 2, 2, 2, 362, 347, 3, 2,
	2, 2, 362, 350, 3, 2, 2, 2, 362, 353, 3, 2, 2, 2, 362, 356, 3, 2, 2, 2,
	363, 366, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365,
	69, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 43, 71, 76, 84, 91, 102, 109, 117,
	123, 138, 142, 155, 161, 172, 181, 186, 189, 194, 199, 203, 211, 215, 222,
	229, 236, 243, 251, 257, 260, 266, 271, 275, 282, 291, 301, 306, 311, 314,
	320, 333, 362, 364,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'::'", "'import'", "'('", "')'", "'return'", "'loop'", "'{'", "'}'",
	"'match'", "'=>'", "','", "'='", "'@'", "'->'", "'impl'", "':'", "'+'",
	"'let'", "'mut'", "'fn'", "'typeForm'", "'trait'", "'^'", "'*'", "'/'",
	"'-'", "'<'", "'>'", "'<='", "'>='", "'!='", "'=='", "'&&'", "'||'", "'?'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"WS", "COMMENT", "ID", "NUM", "STRING",
}

var ruleNames = []string{
	"prog", "topStatList", "topStat", "importMod", "importStat", "statList",
	"stat", "returnStat", "loopStat", "exprStat", "matchRule", "assign", "exprList",
	"fnCall", "typeForm", "typeList", "annotation", "methodList", "method",
	"implBlock", "exportor", "globalVarDefine", "define", "localVarDefine",
	"paramList", "param", "fnDefine", "attrList", "attr", "typeDefine", "tmethodList",
	"tmethod", "traitDefine", "expr",
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
	ElzParserT__24   = 25
	ElzParserT__25   = 26
	ElzParserT__26   = 27
	ElzParserT__27   = 28
	ElzParserT__28   = 29
	ElzParserT__29   = 30
	ElzParserT__30   = 31
	ElzParserT__31   = 32
	ElzParserT__32   = 33
	ElzParserT__33   = 34
	ElzParserT__34   = 35
	ElzParserWS      = 36
	ElzParserCOMMENT = 37
	ElzParserID      = 38
	ElzParserNUM     = 39
	ElzParserSTRING  = 40
)

// ElzParser rules.
const (
	ElzParserRULE_prog            = 0
	ElzParserRULE_topStatList     = 1
	ElzParserRULE_topStat         = 2
	ElzParserRULE_importMod       = 3
	ElzParserRULE_importStat      = 4
	ElzParserRULE_statList        = 5
	ElzParserRULE_stat            = 6
	ElzParserRULE_returnStat      = 7
	ElzParserRULE_loopStat        = 8
	ElzParserRULE_exprStat        = 9
	ElzParserRULE_matchRule       = 10
	ElzParserRULE_assign          = 11
	ElzParserRULE_exprList        = 12
	ElzParserRULE_fnCall          = 13
	ElzParserRULE_typeForm        = 14
	ElzParserRULE_typeList        = 15
	ElzParserRULE_annotation      = 16
	ElzParserRULE_methodList      = 17
	ElzParserRULE_method          = 18
	ElzParserRULE_implBlock       = 19
	ElzParserRULE_exportor        = 20
	ElzParserRULE_globalVarDefine = 21
	ElzParserRULE_define          = 22
	ElzParserRULE_localVarDefine  = 23
	ElzParserRULE_paramList       = 24
	ElzParserRULE_param           = 25
	ElzParserRULE_fnDefine        = 26
	ElzParserRULE_attrList        = 27
	ElzParserRULE_attr            = 28
	ElzParserRULE_typeDefine      = 29
	ElzParserRULE_tmethodList     = 30
	ElzParserRULE_tmethod         = 31
	ElzParserRULE_traitDefine     = 32
	ElzParserRULE_expr            = 33
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__14)|(1<<ElzParserT__16)|(1<<ElzParserT__19)|(1<<ElzParserT__20)|(1<<ElzParserT__21))) != 0) || _la == ElzParserID {
		{
			p.SetState(68)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__14)|(1<<ElzParserT__16)|(1<<ElzParserT__19)|(1<<ElzParserT__20)|(1<<ElzParserT__21))) != 0) || _la == ElzParserID {
		{
			p.SetState(71)
			p.TopStat()
		}

		p.SetState(74)
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

func (s *TopStatContext) GlobalVarDefine() IGlobalVarDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalVarDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalVarDefineContext)
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.FnDefine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.GlobalVarDefine()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.TypeDefine()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.ImplBlock()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)
			p.TraitDefine()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(81)
			p.ImportStat()
		}

	}

	return localctx
}

// IImportModContext is an interface to support dynamic dispatch.
type IImportModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportModContext differentiates from other interfaces.
	IsImportModContext()
}

type ImportModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportModContext() *ImportModContext {
	var p = new(ImportModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_importMod
	return p
}

func (*ImportModContext) IsImportModContext() {}

func NewImportModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportModContext {
	var p = new(ImportModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_importMod

	return p
}

func (s *ImportModContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportModContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ElzParserID)
}

func (s *ImportModContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ElzParserID, i)
}

func (s *ImportModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterImportMod(s)
	}
}

func (s *ImportModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitImportMod(s)
	}
}

func (p *ElzParser) ImportMod() (localctx IImportModContext) {
	localctx = NewImportModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ElzParserRULE_importMod)
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
		p.SetState(84)
		p.Match(ElzParserID)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__0 {
		{
			p.SetState(85)
			p.Match(ElzParserT__0)
		}
		{
			p.SetState(86)
			p.Match(ElzParserID)
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *ImportStatContext) ImportMod() IImportModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportModContext)
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
	p.EnterRule(localctx, 8, ElzParserRULE_importStat)

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
		p.SetState(92)
		p.Match(ElzParserT__1)
	}
	{
		p.SetState(93)
		p.Match(ElzParserT__2)
	}
	{
		p.SetState(94)
		p.ImportMod()
	}
	{
		p.SetState(95)
		p.Match(ElzParserT__3)
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
	p.EnterRule(localctx, 10, ElzParserRULE_statList)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__4)|(1<<ElzParserT__5)|(1<<ElzParserT__8)|(1<<ElzParserT__17))) != 0) || _la == ElzParserID {
		{
			p.SetState(97)
			p.Stat()
		}

		p.SetState(100)
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

func (s *StatContext) LocalVarDefine() ILocalVarDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVarDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalVarDefineContext)
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
	p.EnterRule(localctx, 12, ElzParserRULE_stat)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.LocalVarDefine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.LoopStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.ReturnStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
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
	p.EnterRule(localctx, 14, ElzParserRULE_returnStat)

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
		p.SetState(109)
		p.Match(ElzParserT__4)
	}
	{
		p.SetState(110)
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
	p.EnterRule(localctx, 16, ElzParserRULE_loopStat)
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
		p.SetState(112)
		p.Match(ElzParserT__5)
	}
	{
		p.SetState(113)
		p.Match(ElzParserT__6)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__4)|(1<<ElzParserT__5)|(1<<ElzParserT__8)|(1<<ElzParserT__17))) != 0) || _la == ElzParserID {
		{
			p.SetState(114)
			p.StatList()
		}

	}
	{
		p.SetState(117)
		p.Match(ElzParserT__7)
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
	p.EnterRule(localctx, 18, ElzParserRULE_exprStat)

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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.MatchRule()
		}

	case ElzParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
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
	p.EnterRule(localctx, 20, ElzParserRULE_matchRule)
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
		p.SetState(123)
		p.Match(ElzParserT__8)
	}
	{
		p.SetState(124)
		p.expr(0)
	}
	{
		p.SetState(125)
		p.Match(ElzParserT__6)
	}
	{
		p.SetState(126)
		p.expr(0)
	}
	{
		p.SetState(127)
		p.Match(ElzParserT__9)
	}
	{
		p.SetState(128)
		p.Stat()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(129)
				p.Match(ElzParserT__10)
			}
			{
				p.SetState(130)
				p.expr(0)
			}
			{
				p.SetState(131)
				p.Match(ElzParserT__9)
			}
			{
				p.SetState(132)
				p.Stat()
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__10 {
		{
			p.SetState(139)
			p.Match(ElzParserT__10)
		}

	}
	{
		p.SetState(142)
		p.Match(ElzParserT__7)
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
	p.EnterRule(localctx, 22, ElzParserRULE_assign)

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
		p.SetState(144)
		p.Match(ElzParserID)
	}
	{
		p.SetState(145)
		p.Match(ElzParserT__11)
	}
	{
		p.SetState(146)
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
	p.EnterRule(localctx, 24, ElzParserRULE_exprList)
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
		p.SetState(148)
		p.expr(0)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__10 {
		{
			p.SetState(149)
			p.Match(ElzParserT__10)
		}
		{
			p.SetState(150)
			p.expr(0)
		}

		p.SetState(155)
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
	p.EnterRule(localctx, 26, ElzParserRULE_fnCall)
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
		p.SetState(156)
		p.Match(ElzParserID)
	}
	{
		p.SetState(157)
		p.Match(ElzParserT__2)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__2 || _la == ElzParserT__8 || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ElzParserID-38))|(1<<(ElzParserNUM-38))|(1<<(ElzParserSTRING-38)))) != 0) {
		{
			p.SetState(158)
			p.ExprList()
		}

	}
	{
		p.SetState(161)
		p.Match(ElzParserT__3)
	}

	return localctx
}

// ITypeFormContext is an interface to support dynamic dispatch.
type ITypeFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFormContext differentiates from other interfaces.
	IsTypeFormContext()
}

type TypeFormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFormContext() *TypeFormContext {
	var p = new(TypeFormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_typeForm
	return p
}

func (*TypeFormContext) IsTypeFormContext() {}

func NewTypeFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFormContext {
	var p = new(TypeFormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_typeForm

	return p
}

func (s *TypeFormContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFormContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TypeFormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterTypeForm(s)
	}
}

func (s *TypeFormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitTypeForm(s)
	}
}

func (p *ElzParser) TypeForm() (localctx ITypeFormContext) {
	localctx = NewTypeFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ElzParserRULE_typeForm)

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
		p.SetState(163)
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

func (s *TypeListContext) AllTypeForm() []ITypeFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFormContext)(nil)).Elem())
	var tst = make([]ITypeFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFormContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypeForm(i int) ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 30, ElzParserRULE_typeList)
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
		p.SetState(165)
		p.TypeForm()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__10 {
		{
			p.SetState(166)
			p.Match(ElzParserT__10)
		}
		{
			p.SetState(167)
			p.TypeForm()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *AnnotationContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *ElzParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ElzParserRULE_annotation)
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
		p.SetState(173)
		p.Match(ElzParserT__12)
	}
	{
		p.SetState(174)
		p.Match(ElzParserID)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__2 {
		{
			p.SetState(175)
			p.Match(ElzParserT__2)
		}
		{
			p.SetState(176)
			p.expr(0)
		}
		{
			p.SetState(177)
			p.Match(ElzParserT__3)
		}

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
	p.EnterRule(localctx, 34, ElzParserRULE_methodList)
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
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__16 || _la == ElzParserID {
		{
			p.SetState(181)
			p.Method()
		}

		p.SetState(184)
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

func (s *MethodContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 36, ElzParserRULE_method)
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
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(186)
			p.Exportor()
		}

	}
	{
		p.SetState(189)
		p.Match(ElzParserID)
	}
	{
		p.SetState(190)
		p.Match(ElzParserT__2)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(191)
			p.ParamList()
		}

	}
	{
		p.SetState(194)
		p.Match(ElzParserT__3)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__13 {
		{
			p.SetState(195)
			p.Match(ElzParserT__13)
		}
		{
			p.SetState(196)
			p.TypeForm()
		}

	}
	{
		p.SetState(199)
		p.Match(ElzParserT__6)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__4)|(1<<ElzParserT__5)|(1<<ElzParserT__8)|(1<<ElzParserT__17))) != 0) || _la == ElzParserID {
		{
			p.SetState(200)
			p.StatList()
		}

	}
	{
		p.SetState(203)
		p.Match(ElzParserT__7)
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
	p.EnterRule(localctx, 38, ElzParserRULE_implBlock)
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
		p.SetState(205)
		p.Match(ElzParserT__14)
	}
	{
		p.SetState(206)
		p.Match(ElzParserID)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__15 {
		{
			p.SetState(207)
			p.Match(ElzParserT__15)
		}
		{
			p.SetState(208)
			p.TypeList()
		}

	}
	{
		p.SetState(211)
		p.Match(ElzParserT__6)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 || _la == ElzParserID {
		{
			p.SetState(212)
			p.MethodList()
		}

	}
	{
		p.SetState(215)
		p.Match(ElzParserT__7)
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
	p.EnterRule(localctx, 40, ElzParserRULE_exportor)

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
		p.SetState(217)
		p.Match(ElzParserT__16)
	}

	return localctx
}

// IGlobalVarDefineContext is an interface to support dynamic dispatch.
type IGlobalVarDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalVarDefineContext differentiates from other interfaces.
	IsGlobalVarDefineContext()
}

type GlobalVarDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalVarDefineContext() *GlobalVarDefineContext {
	var p = new(GlobalVarDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_globalVarDefine
	return p
}

func (*GlobalVarDefineContext) IsGlobalVarDefineContext() {}

func NewGlobalVarDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalVarDefineContext {
	var p = new(GlobalVarDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_globalVarDefine

	return p
}

func (s *GlobalVarDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalVarDefineContext) Define() IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *GlobalVarDefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *GlobalVarDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalVarDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalVarDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterGlobalVarDefine(s)
	}
}

func (s *GlobalVarDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitGlobalVarDefine(s)
	}
}

func (p *ElzParser) GlobalVarDefine() (localctx IGlobalVarDefineContext) {
	localctx = NewGlobalVarDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ElzParserRULE_globalVarDefine)
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
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(219)
			p.Exportor()
		}

	}
	{
		p.SetState(222)
		p.Define()
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

func (s *DefineContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 44, ElzParserRULE_define)
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
		p.SetState(224)
		p.Match(ElzParserID)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__15 {
		{
			p.SetState(225)
			p.Match(ElzParserT__15)
		}
		{
			p.SetState(226)
			p.TypeForm()
		}

	}
	{
		p.SetState(229)
		p.Match(ElzParserT__11)
	}
	{
		p.SetState(230)
		p.expr(0)
	}

	return localctx
}

// ILocalVarDefineContext is an interface to support dynamic dispatch.
type ILocalVarDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMut returns the mut token.
	GetMut() antlr.Token

	// SetMut sets the mut token.
	SetMut(antlr.Token)

	// IsLocalVarDefineContext differentiates from other interfaces.
	IsLocalVarDefineContext()
}

type LocalVarDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mut    antlr.Token
}

func NewEmptyLocalVarDefineContext() *LocalVarDefineContext {
	var p = new(LocalVarDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_localVarDefine
	return p
}

func (*LocalVarDefineContext) IsLocalVarDefineContext() {}

func NewLocalVarDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVarDefineContext {
	var p = new(LocalVarDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_localVarDefine

	return p
}

func (s *LocalVarDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVarDefineContext) GetMut() antlr.Token { return s.mut }

func (s *LocalVarDefineContext) SetMut(v antlr.Token) { s.mut = v }

func (s *LocalVarDefineContext) AllDefine() []IDefineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefineContext)(nil)).Elem())
	var tst = make([]IDefineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefineContext)
		}
	}

	return tst
}

func (s *LocalVarDefineContext) Define(i int) IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *LocalVarDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVarDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVarDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterLocalVarDefine(s)
	}
}

func (s *LocalVarDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitLocalVarDefine(s)
	}
}

func (p *ElzParser) LocalVarDefine() (localctx ILocalVarDefineContext) {
	localctx = NewLocalVarDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ElzParserRULE_localVarDefine)
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
		p.SetState(232)
		p.Match(ElzParserT__17)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__18 {
		{
			p.SetState(233)

			var _m = p.Match(ElzParserT__18)

			localctx.(*LocalVarDefineContext).mut = _m
		}

	}
	{
		p.SetState(236)
		p.Define()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(237)
				p.Match(ElzParserT__10)
			}
			{
				p.SetState(238)
				p.Define()
			}

		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, ElzParserRULE_paramList)
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
		p.SetState(244)
		p.Param()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__10 {
		{
			p.SetState(245)
			p.Match(ElzParserT__10)
		}
		{
			p.SetState(246)
			p.Param()
		}

		p.SetState(251)
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

func (s *ParamContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 50, ElzParserRULE_param)
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
		p.SetState(252)
		p.Match(ElzParserID)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__15 {
		{
			p.SetState(253)
			p.Match(ElzParserT__15)
		}
		{
			p.SetState(254)
			p.TypeForm()
		}

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

func (s *FnDefineContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 52, ElzParserRULE_fnDefine)
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
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(257)
			p.Exportor()
		}

	}
	{
		p.SetState(260)
		p.Match(ElzParserT__19)
	}
	{
		p.SetState(261)
		p.Match(ElzParserID)
	}
	{
		p.SetState(262)
		p.Match(ElzParserT__2)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(263)
			p.ParamList()
		}

	}
	{
		p.SetState(266)
		p.Match(ElzParserT__3)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__13 {
		{
			p.SetState(267)
			p.Match(ElzParserT__13)
		}
		{
			p.SetState(268)
			p.TypeForm()
		}

	}
	{
		p.SetState(271)
		p.Match(ElzParserT__6)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__4)|(1<<ElzParserT__5)|(1<<ElzParserT__8)|(1<<ElzParserT__17))) != 0) || _la == ElzParserID {
		{
			p.SetState(272)
			p.StatList()
		}

	}
	{
		p.SetState(275)
		p.Match(ElzParserT__7)
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
	p.EnterRule(localctx, 54, ElzParserRULE_attrList)
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
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__16 {
		{
			p.SetState(277)
			p.Attr()
		}

		p.SetState(280)
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

func (s *AttrContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *AttrContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *AttrContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 56, ElzParserRULE_attr)

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
		p.SetState(282)
		p.Exportor()
	}
	{
		p.SetState(283)
		p.Match(ElzParserID)
	}
	{
		p.SetState(284)
		p.Match(ElzParserT__15)
	}
	{
		p.SetState(285)
		p.TypeForm()
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
	p.EnterRule(localctx, 58, ElzParserRULE_typeDefine)
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
		p.SetState(287)
		p.Match(ElzParserT__20)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(288)
			p.Exportor()
		}

	}
	{
		p.SetState(291)
		p.Match(ElzParserID)
	}
	{
		p.SetState(292)
		p.Match(ElzParserT__2)
	}
	{
		p.SetState(293)
		p.AttrList()
	}
	{
		p.SetState(294)
		p.Match(ElzParserT__3)
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
	p.EnterRule(localctx, 60, ElzParserRULE_tmethodList)
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
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserID {
		{
			p.SetState(296)
			p.Tmethod()
		}

		p.SetState(299)
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

func (s *TmethodContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *TmethodContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
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
	p.EnterRule(localctx, 62, ElzParserRULE_tmethod)
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
		p.SetState(301)
		p.Match(ElzParserID)
	}
	{
		p.SetState(302)
		p.Match(ElzParserT__2)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(303)
			p.TypeList()
		}

	}
	{
		p.SetState(306)
		p.Match(ElzParserT__3)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__13 {
		{
			p.SetState(307)
			p.Match(ElzParserT__13)
		}
		{
			p.SetState(308)
			p.TypeForm()
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

func (s *TraitDefineContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *TraitDefineContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
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
	p.EnterRule(localctx, 64, ElzParserRULE_traitDefine)
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
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(311)
			p.Exportor()
		}

	}
	{
		p.SetState(314)
		p.Match(ElzParserT__21)
	}
	{
		p.SetState(315)
		p.Match(ElzParserID)
	}
	{
		p.SetState(316)
		p.Match(ElzParserT__6)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(317)
			p.TmethodList()
		}

	}
	{
		p.SetState(320)
		p.Match(ElzParserT__7)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotEqContext struct {
	*ExprContext
	op antlr.Token
}

func NewNotEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqContext {
	var p = new(NotEqContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqContext) GetOp() antlr.Token { return s.op }

func (s *NotEqContext) SetOp(v antlr.Token) { s.op = v }

func (s *NotEqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotEqContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterNotEq(s)
	}
}

func (s *NotEqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitNotEq(s)
	}
}

type SubExprContext struct {
	*ExprContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

type NumContext struct {
	*ExprContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) NUM() antlr.TerminalNode {
	return s.GetToken(ElzParserNUM, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitNum(s)
	}
}

type MulOrDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulOrDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulOrDivContext {
	var p = new(MulOrDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulOrDivContext) GetOp() antlr.Token { return s.op }

func (s *MulOrDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulOrDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOrDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulOrDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulOrDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterMulOrDiv(s)
	}
}

func (s *MulOrDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitMulOrDiv(s)
	}
}

type CmpContext struct {
	*ExprContext
	op antlr.Token
}

func NewCmpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpContext {
	var p = new(CmpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CmpContext) GetOp() antlr.Token { return s.op }

func (s *CmpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CmpContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CmpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterCmp(s)
	}
}

func (s *CmpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitCmp(s)
	}
}

type EqContext struct {
	*ExprContext
	op antlr.Token
}

func NewEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqContext {
	var p = new(EqContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqContext) GetOp() antlr.Token { return s.op }

func (s *EqContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitEq(s)
	}
}

type AndOrOrContext struct {
	*ExprContext
	op antlr.Token
}

func NewAndOrOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOrOrContext {
	var p = new(AndOrOrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndOrOrContext) GetOp() antlr.Token { return s.op }

func (s *AndOrOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOrOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOrOrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndOrOrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndOrOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAndOrOr(s)
	}
}

func (s *AndOrOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAndOrOr(s)
	}
}

type ThreeOpCmpContext struct {
	*ExprContext
}

func NewThreeOpCmpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThreeOpCmpContext {
	var p = new(ThreeOpCmpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ThreeOpCmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreeOpCmpContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ThreeOpCmpContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ThreeOpCmpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterThreeOpCmp(s)
	}
}

func (s *ThreeOpCmpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitThreeOpCmp(s)
	}
}

type StrContext struct {
	*ExprContext
}

func NewStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrContext {
	var p = new(StrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) STRING() antlr.TerminalNode {
	return s.GetToken(ElzParserSTRING, 0)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitStr(s)
	}
}

type AddOrSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddOrSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddOrSubContext {
	var p = new(AddOrSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddOrSubContext) GetOp() antlr.Token { return s.op }

func (s *AddOrSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddOrSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOrSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddOrSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddOrSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAddOrSub(s)
	}
}

func (s *AddOrSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAddOrSub(s)
	}
}

type StatExprContext struct {
	*ExprContext
}

func NewStatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatExprContext {
	var p = new(StatExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatExprContext) ExprStat() IExprStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprStatContext)
}

func (s *StatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterStatExpr(s)
	}
}

func (s *StatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitStatExpr(s)
	}
}

type PowContext struct {
	*ExprContext
	op antlr.Token
}

func NewPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowContext {
	var p = new(PowContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowContext) GetOp() antlr.Token { return s.op }

func (s *PowContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitPow(s)
	}
}

type IdContext struct {
	*ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitId(s)
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
	_startState := 66
	p.EnterRecursionRule(localctx, 66, ElzParserRULE_expr, _p)
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
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSubExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(323)
			p.Match(ElzParserT__2)
		}
		{
			p.SetState(324)
			p.expr(0)
		}
		{
			p.SetState(325)
			p.Match(ElzParserT__3)
		}

	case 2:
		localctx = NewStatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(327)
			p.ExprStat()
		}

	case 3:
		localctx = NewNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(328)
			p.Match(ElzParserNUM)
		}

	case 4:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(329)
			p.Match(ElzParserID)
		}

	case 5:
		localctx = NewStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(330)
			p.Match(ElzParserSTRING)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(333)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(334)

					var _m = p.Match(ElzParserT__22)

					localctx.(*PowContext).op = _m
				}
				{
					p.SetState(335)
					p.expr(14)
				}

			case 2:
				localctx = NewMulOrDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(337)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulOrDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__23 || _la == ElzParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulOrDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(338)
					p.expr(13)
				}

			case 3:
				localctx = NewAddOrSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(340)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddOrSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__16 || _la == ElzParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddOrSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(341)
					p.expr(12)
				}

			case 4:
				localctx = NewCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(343)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CmpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__26)|(1<<ElzParserT__27)|(1<<ElzParserT__28)|(1<<ElzParserT__29))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CmpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(344)
					p.expr(11)
				}

			case 5:
				localctx = NewNotEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(346)

					var _m = p.Match(ElzParserT__30)

					localctx.(*NotEqContext).op = _m
				}
				{
					p.SetState(347)
					p.expr(10)
				}

			case 6:
				localctx = NewEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(348)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(349)

					var _m = p.Match(ElzParserT__31)

					localctx.(*EqContext).op = _m
				}
				{
					p.SetState(350)
					p.expr(9)
				}

			case 7:
				localctx = NewAndOrOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(352)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AndOrOrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__32 || _la == ElzParserT__33) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AndOrOrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(353)
					p.expr(8)
				}

			case 8:
				localctx = NewThreeOpCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(355)
					p.Match(ElzParserT__34)
				}
				{
					p.SetState(356)
					p.expr(0)
				}
				{
					p.SetState(357)
					p.Match(ElzParserT__15)
				}
				{
					p.SetState(358)
					p.expr(7)
				}

			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ElzParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 33:
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
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
