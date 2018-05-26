struct null;

struct Label {
  int id;
  char *text;
};

struct Label NewLabel(int id, char *text) {
  struct Label label = {id, text};
  label.id = id;
  label.text = text;

  return label;
}
