import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class App039 extends StatelessWidget {
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text("My first Image")),
        body: Center(
          child: Image.asset('MyImage.png'),
        ),
      ),
    );
  }
}
