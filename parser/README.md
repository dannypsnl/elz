# parser
This is elz's parser part.<br>
It use antlr4 to generated parser.<br>
> Remember antlr's target is Go at here.
## Files
- `Elz.g4` is grammar file.<br>
If you want to contributing this part,
you have to learn [antlr4](http://www.antlr.org/).<br>
The link contains installation, document, examples those you will need.
- `refresh.sh` will generate parser code.<br>
Usage:
```bash
$ bash refresh.sh
```
- `clean.sh` remove all generated code.<br>
Usage:
```bash
$ bash clean.sh
```
