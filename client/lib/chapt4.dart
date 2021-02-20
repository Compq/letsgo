import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class App0402 extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Material(
        child: Center(
          child: Text(highlight("Look at me!")),
        ),
      ),
    );
  }
}

highlight(words) {
  return "*** " + words + " ***";
}
