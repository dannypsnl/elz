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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 434,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 87, 10, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 5, 4, 93, 10, 4, 3, 5, 6, 5, 96, 10, 5, 13, 5,
	14, 5, 97, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 108, 10,
	6, 3, 7, 3, 7, 3, 7, 7, 7, 113, 10, 7, 12, 7, 14, 7, 116, 11, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 124, 10, 9, 13, 9, 14, 9, 125, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 133, 10, 10, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 5, 12, 141, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13,
	147, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 7, 14, 160, 10, 14, 12, 14, 14, 14, 163, 11, 14, 3, 14,
	5, 14, 166, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 7, 16, 177, 10, 16, 12, 16, 14, 16, 180, 11, 16, 3, 17, 3, 17,
	3, 17, 5, 17, 185, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 195, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 200, 10, 19,
	12, 19, 14, 19, 203, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	5, 20, 211, 10, 20, 3, 21, 6, 21, 214, 10, 21, 13, 21, 14, 21, 215, 3,
	22, 5, 22, 219, 10, 22, 3, 22, 3, 22, 3, 22, 5, 22, 224, 10, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 229, 10, 22, 3, 22, 3, 22, 5, 22, 233, 10, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 241, 10, 23, 3, 23, 3, 23,
	5, 23, 245, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 5, 25, 252, 10,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 259, 10, 26, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 5, 27, 266, 10, 27, 3, 27, 3, 27, 3, 27, 7, 27, 271,
	10, 27, 12, 27, 14, 27, 274, 11, 27, 3, 28, 3, 28, 3, 28, 7, 28, 279, 10,
	28, 12, 28, 14, 28, 282, 11, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 5,
	30, 289, 10, 30, 3, 31, 3, 31, 3, 31, 3, 32, 5, 32, 295, 10, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 301, 10, 32, 3, 32, 3, 32, 5, 32, 305, 10,
	32, 3, 32, 3, 32, 5, 32, 309, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 317, 10, 33, 3, 33, 3, 33, 5, 33, 321, 10, 33, 3, 34, 3,
	34, 3, 34, 3, 34, 7, 34, 327, 10, 34, 12, 34, 14, 34, 330, 11, 34, 3, 34,
	3, 34, 3, 35, 6, 35, 335, 10, 35, 13, 35, 14, 35, 336, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 5, 37, 346, 10, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 38, 6, 38, 354, 10, 38, 13, 38, 14, 38, 355, 3, 39, 3,
	39, 3, 39, 5, 39, 361, 10, 39, 3, 39, 3, 39, 3, 39, 5, 39, 366, 10, 39,
	3, 40, 5, 40, 369, 10, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 375, 10,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 5, 41, 389, 10, 41, 3, 41, 3, 41, 5, 41, 393, 10, 41, 3,
	41, 3, 41, 5, 41, 397, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 7, 41, 429, 10, 41, 12, 41, 14, 41, 432, 11, 41, 3,
	41, 2, 3, 80, 42, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 2, 8, 3, 2, 4, 7, 3, 2, 6, 7, 3, 2, 37, 38,
	4, 2, 27, 27, 39, 39, 4, 2, 21, 22, 40, 41, 3, 2, 44, 45, 2, 462, 2, 86,
	3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 92, 3, 2, 2, 2, 8, 95, 3, 2, 2, 2, 10,
	107, 3, 2, 2, 2, 12, 109, 3, 2, 2, 2, 14, 117, 3, 2, 2, 2, 16, 123, 3,
	2, 2, 2, 18, 132, 3, 2, 2, 2, 20, 134, 3, 2, 2, 2, 22, 137, 3, 2, 2, 2,
	24, 146, 3, 2, 2, 2, 26, 148, 3, 2, 2, 2, 28, 169, 3, 2, 2, 2, 30, 173,
	3, 2, 2, 2, 32, 181, 3, 2, 2, 2, 34, 194, 3, 2, 2, 2, 36, 196, 3, 2, 2,
	2, 38, 204, 3, 2, 2, 2, 40, 213, 3, 2, 2, 2, 42, 218, 3, 2, 2, 2, 44, 236,
	3, 2, 2, 2, 46, 248, 3, 2, 2, 2, 48, 251, 3, 2, 2, 2, 50, 255, 3, 2, 2,
	2, 52, 263, 3, 2, 2, 2, 54, 275, 3, 2, 2, 2, 56, 283, 3, 2, 2, 2, 58, 286,
	3, 2, 2, 2, 60, 290, 3, 2, 2, 2, 62, 294, 3, 2, 2, 2, 64, 312, 3, 2, 2,
	2, 66, 322, 3, 2, 2, 2, 68, 334, 3, 2, 2, 2, 70, 338, 3, 2, 2, 2, 72, 343,
	3, 2, 2, 2, 74, 353, 3, 2, 2, 2, 76, 357, 3, 2, 2, 2, 78, 368, 3, 2, 2,
	2, 80, 396, 3, 2, 2, 2, 82, 83, 7, 3, 2, 2, 83, 87, 9, 2, 2, 2, 84, 85,
	7, 8, 2, 2, 85, 87, 9, 3, 2, 2, 86, 82, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2,
	87, 3, 3, 2, 2, 2, 88, 89, 7, 8, 2, 2, 89, 90, 9, 3, 2, 2, 90, 5, 3, 2,
	2, 2, 91, 93, 5, 8, 5, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 7,
	3, 2, 2, 2, 94, 96, 5, 10, 6, 2, 95, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 9, 3, 2, 2, 2, 99, 108, 5,
	62, 32, 2, 100, 108, 5, 66, 34, 2, 101, 108, 5, 48, 25, 2, 102, 108, 5,
	72, 37, 2, 103, 108, 5, 44, 23, 2, 104, 108, 5, 78, 40, 2, 105, 108, 5,
	14, 8, 2, 106, 108, 5, 18, 10, 2, 107, 99, 3, 2, 2, 2, 107, 100, 3, 2,
	2, 2, 107, 101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 107, 103, 3, 2, 2, 2,
	107, 104, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108,
	11, 3, 2, 2, 2, 109, 114, 7, 49, 2, 2, 110, 111, 7, 9, 2, 2, 111, 113,
	7, 49, 2, 2, 112, 110, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 114, 115, 3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	117, 118, 7, 10, 2, 2, 118, 119, 7, 11, 2, 2, 119, 120, 5, 12, 7, 2, 120,
	121, 7, 12, 2, 2, 121, 15, 3, 2, 2, 2, 122, 124, 5, 18, 10, 2, 123, 122,
	3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2,
	2, 2, 126, 17, 3, 2, 2, 2, 127, 133, 5, 52, 27, 2, 128, 133, 5, 22, 12,
	2, 129, 133, 5, 20, 11, 2, 130, 133, 5, 28, 15, 2, 131, 133, 5, 24, 13,
	2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 132,
	130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 19, 3, 2, 2, 2, 134, 135, 7,
	13, 2, 2, 135, 136, 5, 80, 41, 2, 136, 21, 3, 2, 2, 2, 137, 138, 7, 14,
	2, 2, 138, 140, 7, 15, 2, 2, 139, 141, 5, 16, 9, 2, 140, 139, 3, 2, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 16, 2, 2, 143,
	23, 3, 2, 2, 2, 144, 147, 5, 26, 14, 2, 145, 147, 5, 32, 17, 2, 146, 144,
	3, 2, 2, 2, 146, 145, 3, 2, 2, 2, 147, 25, 3, 2, 2, 2, 148, 149, 7, 17,
	2, 2, 149, 150, 5, 80, 41, 2, 150, 151, 7, 15, 2, 2, 151, 152, 5, 80, 41,
	2, 152, 153, 7, 18, 2, 2, 153, 161, 5, 18, 10, 2, 154, 155, 7, 19, 2, 2,
	155, 156, 5, 80, 41, 2, 156, 157, 7, 18, 2, 2, 157, 158, 5, 18, 10, 2,
	158, 160, 3, 2, 2, 2, 159, 154, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161,
	159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161,
	3, 2, 2, 2, 164, 166, 7, 19, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 7, 16, 2, 2, 168, 27, 3, 2, 2, 2,
	169, 170, 7, 49, 2, 2, 170, 171, 7, 20, 2, 2, 171, 172, 5, 80, 41, 2, 172,
	29, 3, 2, 2, 2, 173, 178, 5, 80, 41, 2, 174, 175, 7, 19, 2, 2, 175, 177,
	5, 80, 41, 2, 176, 174, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3,
	2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 31, 3, 2, 2, 2, 180, 178, 3, 2, 2,
	2, 181, 182, 7, 49, 2, 2, 182, 184, 7, 11, 2, 2, 183, 185, 5, 30, 16, 2,
	184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	187, 7, 12, 2, 2, 187, 33, 3, 2, 2, 2, 188, 195, 7, 49, 2, 2, 189, 190,
	7, 49, 2, 2, 190, 191, 7, 21, 2, 2, 191, 192, 5, 36, 19, 2, 192, 193, 7,
	22, 2, 2, 193, 195, 3, 2, 2, 2, 194, 188, 3, 2, 2, 2, 194, 189, 3, 2, 2,
	2, 195, 35, 3, 2, 2, 2, 196, 201, 5, 34, 18, 2, 197, 198, 7, 19, 2, 2,
	198, 200, 5, 34, 18, 2, 199, 197, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201,
	199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 37, 3, 2, 2, 2, 203, 201, 3,
	2, 2, 2, 204, 205, 7, 23, 2, 2, 205, 210, 7, 49, 2, 2, 206, 207, 7, 11,
	2, 2, 207, 208, 5, 80, 41, 2, 208, 209, 7, 12, 2, 2, 209, 211, 3, 2, 2,
	2, 210, 206, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 39, 3, 2, 2, 2, 212,
	214, 5, 42, 22, 2, 213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213,
	3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 41, 3, 2, 2, 2, 217, 219, 5, 46,
	24, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2,
	220, 221, 7, 49, 2, 2, 221, 223, 7, 11, 2, 2, 222, 224, 5, 54, 28, 2, 223,
	222, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 228,
	7, 12, 2, 2, 226, 227, 7, 24, 2, 2, 227, 229, 5, 34, 18, 2, 228, 226, 3,
	2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232, 7, 15, 2,
	2, 231, 233, 5, 16, 9, 2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233,
	234, 3, 2, 2, 2, 234, 235, 7, 16, 2, 2, 235, 43, 3, 2, 2, 2, 236, 237,
	7, 25, 2, 2, 237, 240, 7, 49, 2, 2, 238, 239, 7, 26, 2, 2, 239, 241, 5,
	36, 19, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 3, 2,
	2, 2, 242, 244, 7, 15, 2, 2, 243, 245, 5, 40, 21, 2, 244, 243, 3, 2, 2,
	2, 244, 245, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 7, 16, 2, 2, 247,
	45, 3, 2, 2, 2, 248, 249, 7, 27, 2, 2, 249, 47, 3, 2, 2, 2, 250, 252, 5,
	46, 24, 2, 251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 3, 2,
	2, 2, 253, 254, 5, 50, 26, 2, 254, 49, 3, 2, 2, 2, 255, 258, 7, 49, 2,
	2, 256, 257, 7, 26, 2, 2, 257, 259, 5, 34, 18, 2, 258, 256, 3, 2, 2, 2,
	258, 259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 7, 20, 2, 2, 261,
	262, 5, 80, 41, 2, 262, 51, 3, 2, 2, 2, 263, 265, 7, 28, 2, 2, 264, 266,
	7, 29, 2, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2,
	2, 2, 267, 272, 5, 50, 26, 2, 268, 269, 7, 19, 2, 2, 269, 271, 5, 50, 26,
	2, 270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 53, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 280, 5,
	58, 30, 2, 276, 277, 7, 19, 2, 2, 277, 279, 5, 58, 30, 2, 278, 276, 3,
	2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2,
	2, 281, 55, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 283, 284, 7, 26, 2, 2, 284,
	285, 5, 34, 18, 2, 285, 57, 3, 2, 2, 2, 286, 288, 7, 49, 2, 2, 287, 289,
	5, 56, 29, 2, 288, 287, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 59, 3, 2,
	2, 2, 290, 291, 7, 24, 2, 2, 291, 292, 5, 34, 18, 2, 292, 61, 3, 2, 2,
	2, 293, 295, 5, 46, 24, 2, 294, 293, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2,
	295, 296, 3, 2, 2, 2, 296, 297, 7, 30, 2, 2, 297, 298, 7, 49, 2, 2, 298,
	300, 7, 11, 2, 2, 299, 301, 5, 54, 28, 2, 300, 299, 3, 2, 2, 2, 300, 301,
	3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 304, 7, 12, 2, 2, 303, 305, 5, 60,
	31, 2, 304, 303, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2,
	306, 308, 7, 15, 2, 2, 307, 309, 5, 16, 9, 2, 308, 307, 3, 2, 2, 2, 308,
	309, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 7, 16, 2, 2, 311, 63,
	3, 2, 2, 2, 312, 313, 7, 30, 2, 2, 313, 314, 7, 49, 2, 2, 314, 316, 7,
	11, 2, 2, 315, 317, 5, 36, 19, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2,
	2, 2, 317, 318, 3, 2, 2, 2, 318, 320, 7, 12, 2, 2, 319, 321, 5, 60, 31,
	2, 320, 319, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 65, 3, 2, 2, 2, 322,
	323, 7, 31, 2, 2, 323, 324, 7, 52, 2, 2, 324, 328, 7, 15, 2, 2, 325, 327,
	5, 64, 33, 2, 326, 325, 3, 2, 2, 2, 327, 330, 3, 2, 2, 2, 328, 326, 3,
	2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 331, 3, 2, 2, 2, 330, 328, 3, 2, 2,
	2, 331, 332, 7, 16, 2, 2, 332, 67, 3, 2, 2, 2, 333, 335, 5, 70, 36, 2,
	334, 333, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 336,
	337, 3, 2, 2, 2, 337, 69, 3, 2, 2, 2, 338, 339, 5, 46, 24, 2, 339, 340,
	7, 49, 2, 2, 340, 341, 7, 26, 2, 2, 341, 342, 5, 34, 18, 2, 342, 71, 3,
	2, 2, 2, 343, 345, 7, 32, 2, 2, 344, 346, 5, 46, 24, 2, 345, 344, 3, 2,
	2, 2, 345, 346, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 348, 7, 49, 2, 2,
	348, 349, 7, 11, 2, 2, 349, 350, 5, 68, 35, 2, 350, 351, 7, 12, 2, 2, 351,
	73, 3, 2, 2, 2, 352, 354, 5, 76, 39, 2, 353, 352, 3, 2, 2, 2, 354, 355,
	3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 75, 3, 2,
	2, 2, 357, 358, 7, 49, 2, 2, 358, 360, 7, 11, 2, 2, 359, 361, 5, 36, 19,
	2, 360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362,
	365, 7, 12, 2, 2, 363, 364, 7, 24, 2, 2, 364, 366, 5, 34, 18, 2, 365, 363,
	3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 77, 3, 2, 2, 2, 367, 369, 5, 46,
	24, 2, 368, 367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2,
	370, 371, 7, 33, 2, 2, 371, 372, 7, 49, 2, 2, 372, 374, 7, 15, 2, 2, 373,
	375, 5, 74, 38, 2, 374, 373, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376,
	3, 2, 2, 2, 376, 377, 7, 16, 2, 2, 377, 79, 3, 2, 2, 2, 378, 379, 8, 41,
	1, 2, 379, 380, 7, 34, 2, 2, 380, 397, 5, 80, 41, 18, 381, 382, 7, 11,
	2, 2, 382, 383, 5, 80, 41, 2, 383, 384, 7, 12, 2, 2, 384, 397, 3, 2, 2,
	2, 385, 397, 5, 24, 13, 2, 386, 388, 7, 51, 2, 2, 387, 389, 5, 2, 2, 2,
	388, 387, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 397, 3, 2, 2, 2, 390,
	392, 7, 50, 2, 2, 391, 393, 5, 4, 3, 2, 392, 391, 3, 2, 2, 2, 392, 393,
	3, 2, 2, 2, 393, 397, 3, 2, 2, 2, 394, 397, 7, 49, 2, 2, 395, 397, 7, 52,
	2, 2, 396, 378, 3, 2, 2, 2, 396, 381, 3, 2, 2, 2, 396, 385, 3, 2, 2, 2,
	396, 386, 3, 2, 2, 2, 396, 390, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2, 396,
	395, 3, 2, 2, 2, 397, 430, 3, 2, 2, 2, 398, 399, 12, 16, 2, 2, 399, 400,
	7, 36, 2, 2, 400, 429, 5, 80, 41, 17, 401, 402, 12, 15, 2, 2, 402, 403,
	9, 4, 2, 2, 403, 429, 5, 80, 41, 16, 404, 405, 12, 14, 2, 2, 405, 406,
	9, 5, 2, 2, 406, 429, 5, 80, 41, 15, 407, 408, 12, 13, 2, 2, 408, 409,
	9, 6, 2, 2, 409, 429, 5, 80, 41, 14, 410, 411, 12, 12, 2, 2, 411, 412,
	7, 42, 2, 2, 412, 429, 5, 80, 41, 13, 413, 414, 12, 11, 2, 2, 414, 415,
	7, 43, 2, 2, 415, 429, 5, 80, 41, 12, 416, 417, 12, 10, 2, 2, 417, 418,
	9, 7, 2, 2, 418, 429, 5, 80, 41, 11, 419, 420, 12, 9, 2, 2, 420, 421, 7,
	46, 2, 2, 421, 422, 5, 80, 41, 2, 422, 423, 7, 26, 2, 2, 423, 424, 5, 80,
	41, 10, 424, 429, 3, 2, 2, 2, 425, 426, 12, 17, 2, 2, 426, 427, 7, 35,
	2, 2, 427, 429, 5, 34, 18, 2, 428, 398, 3, 2, 2, 2, 428, 401, 3, 2, 2,
	2, 428, 404, 3, 2, 2, 2, 428, 407, 3, 2, 2, 2, 428, 410, 3, 2, 2, 2, 428,
	413, 3, 2, 2, 2, 428, 416, 3, 2, 2, 2, 428, 419, 3, 2, 2, 2, 428, 425,
	3, 2, 2, 2, 429, 432, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2,
	2, 2, 431, 81, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 50, 86, 92, 97, 107, 114,
	125, 132, 140, 146, 161, 165, 178, 184, 194, 201, 210, 215, 218, 223, 228,
	232, 240, 244, 251, 258, 265, 272, 280, 288, 294, 300, 304, 308, 316, 320,
	328, 336, 345, 355, 360, 365, 368, 374, 388, 392, 396, 428, 430,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "''i'", "'8'", "'16'", "'32'", "'64'", "''f'", "'::'", "'import'",
	"'('", "')'", "'return'", "'loop'", "'{'", "'}'", "'match'", "'=>'", "','",
	"'='", "'<'", "'>'", "'@'", "'->'", "'impl'", "':'", "'+'", "'let'", "'mut'",
	"'fn'", "'extern'", "'typeForm'", "'trait'", "'&'", "'as'", "'^'", "'*'",
	"'/'", "'-'", "'<='", "'>='", "'!='", "'=='", "'&&'", "'||'", "'?'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "WS", "COMMENT", "ID", "FLOAT", "INT",
	"STRING",
}

var ruleNames = []string{
	"intSuffix", "floatSuffix", "prog", "topStatList", "topStat", "importMod",
	"importStat", "statList", "stat", "returnStat", "loopStat", "exprStat",
	"matchRule", "assign", "exprList", "fnCall", "typeForm", "typeList", "annotation",
	"methodList", "method", "implBlock", "exportor", "globalVarDef", "define",
	"localVarDef", "paramList", "paramType", "param", "returnType", "fnDefine",
	"declareFn", "externBlock", "attrList", "attr", "typeDefine", "tmethodList",
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
	ElzParserT__35   = 36
	ElzParserT__36   = 37
	ElzParserT__37   = 38
	ElzParserT__38   = 39
	ElzParserT__39   = 40
	ElzParserT__40   = 41
	ElzParserT__41   = 42
	ElzParserT__42   = 43
	ElzParserT__43   = 44
	ElzParserWS      = 45
	ElzParserCOMMENT = 46
	ElzParserID      = 47
	ElzParserFLOAT   = 48
	ElzParserINT     = 49
	ElzParserSTRING  = 50
)

// ElzParser rules.
const (
	ElzParserRULE_intSuffix    = 0
	ElzParserRULE_floatSuffix  = 1
	ElzParserRULE_prog         = 2
	ElzParserRULE_topStatList  = 3
	ElzParserRULE_topStat      = 4
	ElzParserRULE_importMod    = 5
	ElzParserRULE_importStat   = 6
	ElzParserRULE_statList     = 7
	ElzParserRULE_stat         = 8
	ElzParserRULE_returnStat   = 9
	ElzParserRULE_loopStat     = 10
	ElzParserRULE_exprStat     = 11
	ElzParserRULE_matchRule    = 12
	ElzParserRULE_assign       = 13
	ElzParserRULE_exprList     = 14
	ElzParserRULE_fnCall       = 15
	ElzParserRULE_typeForm     = 16
	ElzParserRULE_typeList     = 17
	ElzParserRULE_annotation   = 18
	ElzParserRULE_methodList   = 19
	ElzParserRULE_method       = 20
	ElzParserRULE_implBlock    = 21
	ElzParserRULE_exportor     = 22
	ElzParserRULE_globalVarDef = 23
	ElzParserRULE_define       = 24
	ElzParserRULE_localVarDef  = 25
	ElzParserRULE_paramList    = 26
	ElzParserRULE_paramType    = 27
	ElzParserRULE_param        = 28
	ElzParserRULE_returnType   = 29
	ElzParserRULE_fnDefine     = 30
	ElzParserRULE_declareFn    = 31
	ElzParserRULE_externBlock  = 32
	ElzParserRULE_attrList     = 33
	ElzParserRULE_attr         = 34
	ElzParserRULE_typeDefine   = 35
	ElzParserRULE_tmethodList  = 36
	ElzParserRULE_tmethod      = 37
	ElzParserRULE_traitDefine  = 38
	ElzParserRULE_expr         = 39
)

// IIntSuffixContext is an interface to support dynamic dispatch.
type IIntSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntSuffixContext differentiates from other interfaces.
	IsIntSuffixContext()
}

type IntSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntSuffixContext() *IntSuffixContext {
	var p = new(IntSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_intSuffix
	return p
}

func (*IntSuffixContext) IsIntSuffixContext() {}

func NewIntSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntSuffixContext {
	var p = new(IntSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_intSuffix

	return p
}

func (s *IntSuffixContext) GetParser() antlr.Parser { return s.parser }
func (s *IntSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterIntSuffix(s)
	}
}

func (s *IntSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitIntSuffix(s)
	}
}

func (p *ElzParser) IntSuffix() (localctx IIntSuffixContext) {
	localctx = NewIntSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ElzParserRULE_intSuffix)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(ElzParserT__0)
		}
		{
			p.SetState(81)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__1)|(1<<ElzParserT__2)|(1<<ElzParserT__3)|(1<<ElzParserT__4))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case ElzParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(ElzParserT__5)
		}
		{
			p.SetState(83)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ElzParserT__3 || _la == ElzParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFloatSuffixContext is an interface to support dynamic dispatch.
type IFloatSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatSuffixContext differentiates from other interfaces.
	IsFloatSuffixContext()
}

type FloatSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatSuffixContext() *FloatSuffixContext {
	var p = new(FloatSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_floatSuffix
	return p
}

func (*FloatSuffixContext) IsFloatSuffixContext() {}

func NewFloatSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatSuffixContext {
	var p = new(FloatSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_floatSuffix

	return p
}

func (s *FloatSuffixContext) GetParser() antlr.Parser { return s.parser }
func (s *FloatSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterFloatSuffix(s)
	}
}

func (s *FloatSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitFloatSuffix(s)
	}
}

func (p *ElzParser) FloatSuffix() (localctx IFloatSuffixContext) {
	localctx = NewFloatSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ElzParserRULE_floatSuffix)
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
		p.SetState(86)
		p.Match(ElzParserT__5)
	}
	{
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ElzParserT__3 || _la == ElzParserT__4) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

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
	p.EnterRule(localctx, 4, ElzParserRULE_prog)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__7)|(1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__22)|(1<<ElzParserT__24)|(1<<ElzParserT__25)|(1<<ElzParserT__27)|(1<<ElzParserT__28)|(1<<ElzParserT__29)|(1<<ElzParserT__30))) != 0) || _la == ElzParserID {
		{
			p.SetState(89)
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
	p.EnterRule(localctx, 6, ElzParserRULE_topStatList)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__7)|(1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__22)|(1<<ElzParserT__24)|(1<<ElzParserT__25)|(1<<ElzParserT__27)|(1<<ElzParserT__28)|(1<<ElzParserT__29)|(1<<ElzParserT__30))) != 0) || _la == ElzParserID {
		{
			p.SetState(92)
			p.TopStat()
		}

		p.SetState(95)
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

func (s *TopStatContext) ExternBlock() IExternBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternBlockContext)
}

func (s *TopStatContext) GlobalVarDef() IGlobalVarDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalVarDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobalVarDefContext)
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

func (s *TopStatContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
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
	p.EnterRule(localctx, 8, ElzParserRULE_topStat)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.FnDefine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.ExternBlock()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.GlobalVarDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.TypeDefine()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(101)
			p.ImplBlock()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)
			p.TraitDefine()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(103)
			p.ImportStat()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(104)
			p.Stat()
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
	p.EnterRule(localctx, 10, ElzParserRULE_importMod)
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
		p.SetState(107)
		p.Match(ElzParserID)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__6 {
		{
			p.SetState(108)
			p.Match(ElzParserT__6)
		}
		{
			p.SetState(109)
			p.Match(ElzParserID)
		}

		p.SetState(114)
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
	p.EnterRule(localctx, 12, ElzParserRULE_importStat)

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
		p.SetState(115)
		p.Match(ElzParserT__7)
	}
	{
		p.SetState(116)
		p.Match(ElzParserT__8)
	}
	{
		p.SetState(117)
		p.ImportMod()
	}
	{
		p.SetState(118)
		p.Match(ElzParserT__9)
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
	p.EnterRule(localctx, 14, ElzParserRULE_statList)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__25))) != 0) || _la == ElzParserID {
		{
			p.SetState(120)
			p.Stat()
		}

		p.SetState(123)
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

func (s *StatContext) LocalVarDef() ILocalVarDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVarDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalVarDefContext)
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
	p.EnterRule(localctx, 16, ElzParserRULE_stat)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.LocalVarDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.LoopStat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.ReturnStat()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(128)
			p.Assign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(129)
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
	p.EnterRule(localctx, 18, ElzParserRULE_returnStat)

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
		p.SetState(132)
		p.Match(ElzParserT__10)
	}
	{
		p.SetState(133)
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
	p.EnterRule(localctx, 20, ElzParserRULE_loopStat)
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
		p.SetState(135)
		p.Match(ElzParserT__11)
	}
	{
		p.SetState(136)
		p.Match(ElzParserT__12)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__25))) != 0) || _la == ElzParserID {
		{
			p.SetState(137)
			p.StatList()
		}

	}
	{
		p.SetState(140)
		p.Match(ElzParserT__13)
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
	p.EnterRule(localctx, 22, ElzParserRULE_exprStat)

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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ElzParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.MatchRule()
		}

	case ElzParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
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
	p.EnterRule(localctx, 24, ElzParserRULE_matchRule)
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
		p.SetState(146)
		p.Match(ElzParserT__14)
	}
	{
		p.SetState(147)
		p.expr(0)
	}
	{
		p.SetState(148)
		p.Match(ElzParserT__12)
	}
	{
		p.SetState(149)
		p.expr(0)
	}
	{
		p.SetState(150)
		p.Match(ElzParserT__15)
	}
	{
		p.SetState(151)
		p.Stat()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(152)
				p.Match(ElzParserT__16)
			}
			{
				p.SetState(153)
				p.expr(0)
			}
			{
				p.SetState(154)
				p.Match(ElzParserT__15)
			}
			{
				p.SetState(155)
				p.Stat()
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__16 {
		{
			p.SetState(162)
			p.Match(ElzParserT__16)
		}

	}
	{
		p.SetState(165)
		p.Match(ElzParserT__13)
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
	p.EnterRule(localctx, 26, ElzParserRULE_assign)

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
		p.SetState(167)
		p.Match(ElzParserID)
	}
	{
		p.SetState(168)
		p.Match(ElzParserT__17)
	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 28, ElzParserRULE_exprList)
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
		p.SetState(171)
		p.expr(0)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__16 {
		{
			p.SetState(172)
			p.Match(ElzParserT__16)
		}
		{
			p.SetState(173)
			p.expr(0)
		}

		p.SetState(178)
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
	p.EnterRule(localctx, 30, ElzParserRULE_fnCall)
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
		p.SetState(179)
		p.Match(ElzParserID)
	}
	{
		p.SetState(180)
		p.Match(ElzParserT__8)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__8 || _la == ElzParserT__14 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ElzParserT__31-32))|(1<<(ElzParserID-32))|(1<<(ElzParserFLOAT-32))|(1<<(ElzParserINT-32))|(1<<(ElzParserSTRING-32)))) != 0) {
		{
			p.SetState(181)
			p.ExprList()
		}

	}
	{
		p.SetState(184)
		p.Match(ElzParserT__9)
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

func (s *TypeFormContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
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
	p.EnterRule(localctx, 32, ElzParserRULE_typeForm)

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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(ElzParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Match(ElzParserID)
		}
		{
			p.SetState(188)
			p.Match(ElzParserT__18)
		}
		{
			p.SetState(189)
			p.TypeList()
		}
		{
			p.SetState(190)
			p.Match(ElzParserT__19)
		}

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
	p.EnterRule(localctx, 34, ElzParserRULE_typeList)
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
		p.SetState(194)
		p.TypeForm()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__16 {
		{
			p.SetState(195)
			p.Match(ElzParserT__16)
		}
		{
			p.SetState(196)
			p.TypeForm()
		}

		p.SetState(201)
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
	p.EnterRule(localctx, 36, ElzParserRULE_annotation)
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
		p.SetState(202)
		p.Match(ElzParserT__20)
	}
	{
		p.SetState(203)
		p.Match(ElzParserID)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__8 {
		{
			p.SetState(204)
			p.Match(ElzParserT__8)
		}
		{
			p.SetState(205)
			p.expr(0)
		}
		{
			p.SetState(206)
			p.Match(ElzParserT__9)
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
	p.EnterRule(localctx, 38, ElzParserRULE_methodList)
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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__24 || _la == ElzParserID {
		{
			p.SetState(210)
			p.Method()
		}

		p.SetState(213)
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
	p.EnterRule(localctx, 40, ElzParserRULE_method)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 {
		{
			p.SetState(215)
			p.Exportor()
		}

	}
	{
		p.SetState(218)
		p.Match(ElzParserID)
	}
	{
		p.SetState(219)
		p.Match(ElzParserT__8)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(220)
			p.ParamList()
		}

	}
	{
		p.SetState(223)
		p.Match(ElzParserT__9)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__21 {
		{
			p.SetState(224)
			p.Match(ElzParserT__21)
		}
		{
			p.SetState(225)
			p.TypeForm()
		}

	}
	{
		p.SetState(228)
		p.Match(ElzParserT__12)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__25))) != 0) || _la == ElzParserID {
		{
			p.SetState(229)
			p.StatList()
		}

	}
	{
		p.SetState(232)
		p.Match(ElzParserT__13)
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
	p.EnterRule(localctx, 42, ElzParserRULE_implBlock)
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
		p.SetState(234)
		p.Match(ElzParserT__22)
	}
	{
		p.SetState(235)
		p.Match(ElzParserID)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__23 {
		{
			p.SetState(236)
			p.Match(ElzParserT__23)
		}
		{
			p.SetState(237)
			p.TypeList()
		}

	}
	{
		p.SetState(240)
		p.Match(ElzParserT__12)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 || _la == ElzParserID {
		{
			p.SetState(241)
			p.MethodList()
		}

	}
	{
		p.SetState(244)
		p.Match(ElzParserT__13)
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
	p.EnterRule(localctx, 44, ElzParserRULE_exportor)

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
		p.SetState(246)
		p.Match(ElzParserT__24)
	}

	return localctx
}

// IGlobalVarDefContext is an interface to support dynamic dispatch.
type IGlobalVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalVarDefContext differentiates from other interfaces.
	IsGlobalVarDefContext()
}

type GlobalVarDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalVarDefContext() *GlobalVarDefContext {
	var p = new(GlobalVarDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_globalVarDef
	return p
}

func (*GlobalVarDefContext) IsGlobalVarDefContext() {}

func NewGlobalVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalVarDefContext {
	var p = new(GlobalVarDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_globalVarDef

	return p
}

func (s *GlobalVarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalVarDefContext) Define() IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *GlobalVarDefContext) Exportor() IExportorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportorContext)
}

func (s *GlobalVarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalVarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalVarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterGlobalVarDef(s)
	}
}

func (s *GlobalVarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitGlobalVarDef(s)
	}
}

func (p *ElzParser) GlobalVarDef() (localctx IGlobalVarDefContext) {
	localctx = NewGlobalVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ElzParserRULE_globalVarDef)
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
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 {
		{
			p.SetState(248)
			p.Exportor()
		}

	}
	{
		p.SetState(251)
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
	p.EnterRule(localctx, 48, ElzParserRULE_define)
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
		p.SetState(253)
		p.Match(ElzParserID)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__23 {
		{
			p.SetState(254)
			p.Match(ElzParserT__23)
		}
		{
			p.SetState(255)
			p.TypeForm()
		}

	}
	{
		p.SetState(258)
		p.Match(ElzParserT__17)
	}
	{
		p.SetState(259)
		p.expr(0)
	}

	return localctx
}

// ILocalVarDefContext is an interface to support dynamic dispatch.
type ILocalVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMut returns the mut token.
	GetMut() antlr.Token

	// SetMut sets the mut token.
	SetMut(antlr.Token)

	// IsLocalVarDefContext differentiates from other interfaces.
	IsLocalVarDefContext()
}

type LocalVarDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mut    antlr.Token
}

func NewEmptyLocalVarDefContext() *LocalVarDefContext {
	var p = new(LocalVarDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_localVarDef
	return p
}

func (*LocalVarDefContext) IsLocalVarDefContext() {}

func NewLocalVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVarDefContext {
	var p = new(LocalVarDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_localVarDef

	return p
}

func (s *LocalVarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVarDefContext) GetMut() antlr.Token { return s.mut }

func (s *LocalVarDefContext) SetMut(v antlr.Token) { s.mut = v }

func (s *LocalVarDefContext) AllDefine() []IDefineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefineContext)(nil)).Elem())
	var tst = make([]IDefineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefineContext)
		}
	}

	return tst
}

func (s *LocalVarDefContext) Define(i int) IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *LocalVarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterLocalVarDef(s)
	}
}

func (s *LocalVarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitLocalVarDef(s)
	}
}

func (p *ElzParser) LocalVarDef() (localctx ILocalVarDefContext) {
	localctx = NewLocalVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ElzParserRULE_localVarDef)
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
		p.SetState(261)
		p.Match(ElzParserT__25)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__26 {
		{
			p.SetState(262)

			var _m = p.Match(ElzParserT__26)

			localctx.(*LocalVarDefContext).mut = _m
		}

	}
	{
		p.SetState(265)
		p.Define()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(266)
				p.Match(ElzParserT__16)
			}
			{
				p.SetState(267)
				p.Define()
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, ElzParserRULE_paramList)
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
		p.SetState(273)
		p.Param()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__16 {
		{
			p.SetState(274)
			p.Match(ElzParserT__16)
		}
		{
			p.SetState(275)
			p.Param()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamTypeContext is an interface to support dynamic dispatch.
type IParamTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamTypeContext differentiates from other interfaces.
	IsParamTypeContext()
}

type ParamTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamTypeContext() *ParamTypeContext {
	var p = new(ParamTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_paramType
	return p
}

func (*ParamTypeContext) IsParamTypeContext() {}

func NewParamTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamTypeContext {
	var p = new(ParamTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_paramType

	return p
}

func (s *ParamTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamTypeContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
}

func (s *ParamTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterParamType(s)
	}
}

func (s *ParamTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitParamType(s)
	}
}

func (p *ElzParser) ParamType() (localctx IParamTypeContext) {
	localctx = NewParamTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ElzParserRULE_paramType)

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
		p.SetState(281)
		p.Match(ElzParserT__23)
	}
	{
		p.SetState(282)
		p.TypeForm()
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

func (s *ParamContext) ParamType() IParamTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
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
	p.EnterRule(localctx, 56, ElzParserRULE_param)
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
		p.SetState(284)
		p.Match(ElzParserID)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__23 {
		{
			p.SetState(285)
			p.ParamType()
		}

	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *ElzParser) ReturnType() (localctx IReturnTypeContext) {
	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ElzParserRULE_returnType)

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
		p.Match(ElzParserT__21)
	}
	{
		p.SetState(289)
		p.TypeForm()
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

func (s *FnDefineContext) ReturnType() IReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
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
	p.EnterRule(localctx, 60, ElzParserRULE_fnDefine)
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
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 {
		{
			p.SetState(291)
			p.Exportor()
		}

	}
	{
		p.SetState(294)
		p.Match(ElzParserT__27)
	}
	{
		p.SetState(295)
		p.Match(ElzParserID)
	}
	{
		p.SetState(296)
		p.Match(ElzParserT__8)
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(297)
			p.ParamList()
		}

	}
	{
		p.SetState(300)
		p.Match(ElzParserT__9)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__21 {
		{
			p.SetState(301)
			p.ReturnType()
		}

	}
	{
		p.SetState(304)
		p.Match(ElzParserT__12)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ElzParserT__10)|(1<<ElzParserT__11)|(1<<ElzParserT__14)|(1<<ElzParserT__25))) != 0) || _la == ElzParserID {
		{
			p.SetState(305)
			p.StatList()
		}

	}
	{
		p.SetState(308)
		p.Match(ElzParserT__13)
	}

	return localctx
}

// IDeclareFnContext is an interface to support dynamic dispatch.
type IDeclareFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclareFnContext differentiates from other interfaces.
	IsDeclareFnContext()
}

type DeclareFnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclareFnContext() *DeclareFnContext {
	var p = new(DeclareFnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_declareFn
	return p
}

func (*DeclareFnContext) IsDeclareFnContext() {}

func NewDeclareFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclareFnContext {
	var p = new(DeclareFnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_declareFn

	return p
}

func (s *DeclareFnContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclareFnContext) ID() antlr.TerminalNode {
	return s.GetToken(ElzParserID, 0)
}

func (s *DeclareFnContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *DeclareFnContext) ReturnType() IReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *DeclareFnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclareFnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclareFnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterDeclareFn(s)
	}
}

func (s *DeclareFnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitDeclareFn(s)
	}
}

func (p *ElzParser) DeclareFn() (localctx IDeclareFnContext) {
	localctx = NewDeclareFnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ElzParserRULE_declareFn)
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
		p.SetState(310)
		p.Match(ElzParserT__27)
	}
	{
		p.SetState(311)
		p.Match(ElzParserID)
	}
	{
		p.SetState(312)
		p.Match(ElzParserT__8)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(313)
			p.TypeList()
		}

	}
	{
		p.SetState(316)
		p.Match(ElzParserT__9)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__21 {
		{
			p.SetState(317)
			p.ReturnType()
		}

	}

	return localctx
}

// IExternBlockContext is an interface to support dynamic dispatch.
type IExternBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternBlockContext differentiates from other interfaces.
	IsExternBlockContext()
}

type ExternBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternBlockContext() *ExternBlockContext {
	var p = new(ExternBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElzParserRULE_externBlock
	return p
}

func (*ExternBlockContext) IsExternBlockContext() {}

func NewExternBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternBlockContext {
	var p = new(ExternBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElzParserRULE_externBlock

	return p
}

func (s *ExternBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternBlockContext) STRING() antlr.TerminalNode {
	return s.GetToken(ElzParserSTRING, 0)
}

func (s *ExternBlockContext) AllDeclareFn() []IDeclareFnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclareFnContext)(nil)).Elem())
	var tst = make([]IDeclareFnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclareFnContext)
		}
	}

	return tst
}

func (s *ExternBlockContext) DeclareFn(i int) IDeclareFnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclareFnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclareFnContext)
}

func (s *ExternBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterExternBlock(s)
	}
}

func (s *ExternBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitExternBlock(s)
	}
}

func (p *ElzParser) ExternBlock() (localctx IExternBlockContext) {
	localctx = NewExternBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ElzParserRULE_externBlock)
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
		p.SetState(320)
		p.Match(ElzParserT__28)
	}
	{
		p.SetState(321)
		p.Match(ElzParserSTRING)
	}
	{
		p.SetState(322)
		p.Match(ElzParserT__12)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ElzParserT__27 {
		{
			p.SetState(323)
			p.DeclareFn()
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(329)
		p.Match(ElzParserT__13)
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
	p.EnterRule(localctx, 66, ElzParserRULE_attrList)
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
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserT__24 {
		{
			p.SetState(331)
			p.Attr()
		}

		p.SetState(334)
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
	p.EnterRule(localctx, 68, ElzParserRULE_attr)

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
		p.SetState(336)
		p.Exportor()
	}
	{
		p.SetState(337)
		p.Match(ElzParserID)
	}
	{
		p.SetState(338)
		p.Match(ElzParserT__23)
	}
	{
		p.SetState(339)
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
	p.EnterRule(localctx, 70, ElzParserRULE_typeDefine)
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
		p.SetState(341)
		p.Match(ElzParserT__29)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 {
		{
			p.SetState(342)
			p.Exportor()
		}

	}
	{
		p.SetState(345)
		p.Match(ElzParserID)
	}
	{
		p.SetState(346)
		p.Match(ElzParserT__8)
	}
	{
		p.SetState(347)
		p.AttrList()
	}
	{
		p.SetState(348)
		p.Match(ElzParserT__9)
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
	p.EnterRule(localctx, 72, ElzParserRULE_tmethodList)
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
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ElzParserID {
		{
			p.SetState(350)
			p.Tmethod()
		}

		p.SetState(353)
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
	p.EnterRule(localctx, 74, ElzParserRULE_tmethod)
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
		p.SetState(355)
		p.Match(ElzParserID)
	}
	{
		p.SetState(356)
		p.Match(ElzParserT__8)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(357)
			p.TypeList()
		}

	}
	{
		p.SetState(360)
		p.Match(ElzParserT__9)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__21 {
		{
			p.SetState(361)
			p.Match(ElzParserT__21)
		}
		{
			p.SetState(362)
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
	p.EnterRule(localctx, 76, ElzParserRULE_traitDefine)
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
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserT__24 {
		{
			p.SetState(365)
			p.Exportor()
		}

	}
	{
		p.SetState(368)
		p.Match(ElzParserT__30)
	}
	{
		p.SetState(369)
		p.Match(ElzParserID)
	}
	{
		p.SetState(370)
		p.Match(ElzParserT__12)
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ElzParserID {
		{
			p.SetState(371)
			p.TmethodList()
		}

	}
	{
		p.SetState(374)
		p.Match(ElzParserT__13)
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

type IntContext struct {
	*ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(ElzParserINT, 0)
}

func (s *IntContext) IntSuffix() IIntSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntSuffixContext)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitInt(s)
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

type FloatContext struct {
	*ExprContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ElzParserFLOAT, 0)
}

func (s *FloatContext) FloatSuffix() IFloatSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatSuffixContext)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitFloat(s)
	}
}

type RefContext struct {
	*ExprContext
	op antlr.Token
}

func NewRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefContext {
	var p = new(RefContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RefContext) GetOp() antlr.Token { return s.op }

func (s *RefContext) SetOp(v antlr.Token) { s.op = v }

func (s *RefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterRef(s)
	}
}

func (s *RefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitRef(s)
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

type AsContext struct {
	*ExprContext
	op antlr.Token
}

func NewAsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsContext {
	var p = new(AsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AsContext) GetOp() antlr.Token { return s.op }

func (s *AsContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
}

func (s *AsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.EnterAs(s)
	}
}

func (s *AsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElzListener); ok {
		listenerT.ExitAs(s)
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
	_startState := 78
	p.EnterRecursionRule(localctx, 78, ElzParserRULE_expr, _p)
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
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRefContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(377)

			var _m = p.Match(ElzParserT__31)

			localctx.(*RefContext).op = _m
		}
		{
			p.SetState(378)
			p.expr(16)
		}

	case 2:
		localctx = NewSubExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(379)
			p.Match(ElzParserT__8)
		}
		{
			p.SetState(380)
			p.expr(0)
		}
		{
			p.SetState(381)
			p.Match(ElzParserT__9)
		}

	case 3:
		localctx = NewStatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(383)
			p.ExprStat()
		}

	case 4:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(384)
			p.Match(ElzParserINT)
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(385)
				p.IntSuffix()
			}

		}

	case 5:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(388)
			p.Match(ElzParserFLOAT)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(389)
				p.FloatSuffix()
			}

		}

	case 6:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(392)
			p.Match(ElzParserID)
		}

	case 7:
		localctx = NewStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)
			p.Match(ElzParserSTRING)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(397)

					var _m = p.Match(ElzParserT__33)

					localctx.(*PowContext).op = _m
				}
				{
					p.SetState(398)
					p.expr(15)
				}

			case 2:
				localctx = NewMulOrDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(399)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(400)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulOrDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__34 || _la == ElzParserT__35) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulOrDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(401)
					p.expr(14)
				}

			case 3:
				localctx = NewAddOrSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(403)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddOrSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__24 || _la == ElzParserT__36) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddOrSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(404)
					p.expr(13)
				}

			case 4:
				localctx = NewCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(406)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CmpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(ElzParserT__18-19))|(1<<(ElzParserT__19-19))|(1<<(ElzParserT__37-19))|(1<<(ElzParserT__38-19)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CmpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(407)
					p.expr(12)
				}

			case 5:
				localctx = NewNotEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(409)

					var _m = p.Match(ElzParserT__39)

					localctx.(*NotEqContext).op = _m
				}
				{
					p.SetState(410)
					p.expr(11)
				}

			case 6:
				localctx = NewEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(412)

					var _m = p.Match(ElzParserT__40)

					localctx.(*EqContext).op = _m
				}
				{
					p.SetState(413)
					p.expr(10)
				}

			case 7:
				localctx = NewAndOrOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(415)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AndOrOrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ElzParserT__41 || _la == ElzParserT__42) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AndOrOrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(416)
					p.expr(9)
				}

			case 8:
				localctx = NewThreeOpCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(418)
					p.Match(ElzParserT__43)
				}
				{
					p.SetState(419)
					p.expr(0)
				}
				{
					p.SetState(420)
					p.Match(ElzParserT__23)
				}
				{
					p.SetState(421)
					p.expr(8)
				}

			case 9:
				localctx = NewAsContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ElzParserRULE_expr)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(424)

					var _m = p.Match(ElzParserT__32)

					localctx.(*AsContext).op = _m
				}
				{
					p.SetState(425)
					p.TypeForm()
				}

			}

		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ElzParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 39:
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
