import 'package:flutter/material.dart';

class GooglePixel3XL1 extends StatelessWidget {
  final ImageProvider phone2;
  GooglePixel3XL1({
    Key key,
    this.phone2 = const AssetImage('assets/images/phone2.png'),
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xffffffff),
      body: Stack(
        children: <Widget>[
          Transform.translate(
            offset: Offset(66.0, 117.0),
            child: Container(
              width: 800.0,
              height: 600.0,
              decoration: BoxDecoration(
                image: DecorationImage(
                  image: phone2,
                  fit: BoxFit.fill,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
