// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Lamees Hemdan
// Created on: March 2023
// This file contains the JS functions for index.html

// This function calculates the area of a trapezoid

function calculateArea () {
  // input
  const base1 = parseInt(document.getElementById('base1').value)
  const base2 = parseInt(document.getElementById('base2').value)
  const height = parseInt(document.getElementById('height').value)

  // process
  const area = (base1 + base2) / 2 * height

  // output
  document.getElementById('area').innerHTML = 'The Area is ' + area + ' cm²'
}
