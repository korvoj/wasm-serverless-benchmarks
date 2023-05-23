use hound;
use std::f32::consts::PI;
use std::i16;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() == 2 {
      let spec = hound::WavSpec {
          channels: 1,
          sample_rate: 44100,
          bits_per_sample: 16,
          sample_format: hound::SampleFormat::Int,
      };
      let mut writer = hound::WavWriter::create(&args[1], spec).unwrap();
      for t in (0..44100).map(|x| x as f32 / 44100.0) {
          let sample = (t * 440.0 * 2.0 * PI).sin();
          let amplitude = i16::MAX as f32;
          writer.write_sample((sample * amplitude) as i16).unwrap();
      }
    }
    else {
      panic!("Pass path for destination .wav file as first argument...");
    }
}
