export default class Workout {
  constructor() {
    this.id = 313;
    this.programId = "";
    this.start = new Date();
    this.end = new Date();
    this.end.setHours(this.start.getHours()+1);
    this.notes = "Foo seems to be bar here?";
    this.sets = [];
  }
}