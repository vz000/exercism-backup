// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let total = 0;
  for (let i = 0; i < birdsPerDay.length; i++) {
    total += birdsPerDay[i];
  }
  return total;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  const daysPerWeek = 7;
  let weekInitializer = 0;
  if (week > 1) {
    weekInitializer = daysPerWeek * (week - 1);
  }
  const newBirdPerDay = [];
  for(let i = weekInitializer; i < (weekInitializer + daysPerWeek); i++) {
    newBirdPerDay.push(birdsPerDay[i]);
  }
  return totalBirdCount(newBirdPerDay);
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {void} should not return anything
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let i = 0; i < birdsPerDay.length; i+= 2) {
    birdsPerDay[i] += 1;
  }
}
