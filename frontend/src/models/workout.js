export default class Workout {
  constructor() {
    // Unique identifier of the workout created by the backend.
    this.id = 313;
    // Optional program id referring to existing program.
    this.programId = "";
    // The time when the workout started.
    this.start = new Date();
    // The time when the workout ended.
    this.end = new Date();
    this.end.setHours(this.start.getHours()+1);
    // Additional user notes about the workout.
    this.notes = "Foo seems to be bar here?";
    // Sets within the workout.
    this.sets = [];
  }
}