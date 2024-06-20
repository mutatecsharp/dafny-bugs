// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(globalState) {
      return (new BigNumber(529)).plus((new BigNumber(64)).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length))).length), new BigNumber(695))).cardinality()), new BigNumber(768))).length)));
    };
    static fm1(p0, p1, p2, globalState) {
      return (((false) ? (true) : (true))) || (true);
    };
    static fm2(globalState) {
      return ((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.FromArray((_dafny.Seq.of(true))), _dafny.MultiSet.fromElements(true))).length),_module.D0.create_DC2(true)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), function (_0_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length))).Keys.Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.FromArray((_dafny.Seq.of(true))), _dafny.MultiSet.fromElements(true))).length),_module.D0.create_DC2(true)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), function (_0_i0) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length))).contains(_1_v0)) {
            _coll0.push([_1_v0,new BigNumber(329)]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),_module.D0.create_DC2(false)),new BigNumber((_dafny.Seq.of(false, true, false, false, false)).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-222),_module.D0.create_DC2(true)),new BigNumber((_dafny.Seq.UnicodeFromString("swxgyamao")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(679))).cardinality()),_module.D0.create_DC2(true)),new BigNumber((_dafny.Seq.of(new BigNumber(273), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(511))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(127),!(false))).length))).length))));
    };
    static fm3(p0, globalState) {
      return _module.D0.create_DC2(!(false) || (!(true)));
    };
    static fm4(globalState) {
      return _dafny.Seq.Concat((_dafny.Seq.of(!(!(false)))), _dafny.Seq.of(true, true));
    };
    static fm5(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-864),new BigNumber(937))).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)),new BigNumber(311))).length)));
    };
    static fm8(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hxltqisx"), _dafny.Seq.UnicodeFromString("tbiom"));
    };
    static fm9(p0, p1, globalState) {
      if (false) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }
    };
    static fm10(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC0(new BigNumber(566)), _module.D0.create_DC0(new BigNumber(607))), _dafny.Seq.of(_module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('g'.codePointAt(0)))).length)), _module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(true, false, false, false, !(false))).length))));
    };
    static fm11(p0, globalState) {
      return function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bb"), _dafny.Seq.UnicodeFromString("mwmmnrp"), _dafny.Seq.UnicodeFromString("obby"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), function (_2_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }))).Elements) {
          let _3_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bb"), _dafny.Seq.UnicodeFromString("mwmmnrp"), _dafny.Seq.UnicodeFromString("obby"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), function (_2_i0) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })), _3_v0)) {
            _coll1.push([_3_v0,(new BigNumber(405)).plus(new BigNumber(921))]);
          }
        }
        return _coll1;
      }();
    };
    static fm12(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(true))).Union(_dafny.Set.fromElements(!(true), false, false));
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-131)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("bqfasi")).length)),!(true) || (!(true)));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements((true) === (false));
    };
    static fm15(p0, p1, p2, globalState) {
      if (!(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(771),false)).contains(new BigNumber(531))) {
        return function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(942)), function (_4_i0) {
            return new BigNumber(-874);
          })).Elements) {
            let _5_v0 = _compr_2;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(942)), function (_4_i0) {
              return new BigNumber(-874);
            }), _5_v0)) {
              _coll2.push([(_5_v0).multipliedBy(new BigNumber(871)),new BigNumber((function () {
                let _coll3 = new _dafny.Map();
                for (const _compr_3 of ((_module.D8.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(760),function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(976), new BigNumber(745))) {
    let _6_v2 = _compr_4;
    if (((new BigNumber(976)).isLessThanOrEqualTo(_6_v2)) && ((_6_v2).isLessThan(new BigNumber(745)))) {
      _coll4.push([_module.__default.safeModuloInt(_6_v2, new BigNumber(710)),new BigNumber((_dafny.Seq.of(!(false), false)).length)]);
    }
  }
  return _coll4;
}()))).dtor_cf32).Keys.Elements) {
                  let _7_v1 = _compr_3;
                  if (((_module.D8.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(760),function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(976), new BigNumber(745))) {
    let _6_v2 = _compr_5;
    if (((new BigNumber(976)).isLessThanOrEqualTo(_6_v2)) && ((_6_v2).isLessThan(new BigNumber(745)))) {
      _coll5.push([_module.__default.safeModuloInt(_6_v2, new BigNumber(710)),new BigNumber((_dafny.Seq.of(!(false), false)).length)]);
    }
  }
  return _coll5;
}()))).dtor_cf32).contains(_7_v1)) {
                    _coll3.push([(_7_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
                      let _coll6 = new _dafny.Set();
                      for (const _compr_6 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))).Elements) {
                        let _8_v3 = _compr_6;
                        if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))).contains(_8_v3)) {
                          _coll6.add(_8_v3);
                        }
                      }
                      return _coll6;
                    }()).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), function (_9_i1) {
                      return new BigNumber((_dafny.Seq.UnicodeFromString("mamhl")).length);
                    })).length))).length)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),true)]);
                  }
                }
                return _coll3;
              }()).length)]);
            }
          }
          return _coll2;
        }();
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(95),new BigNumber(-232))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(967))).length),new BigNumber(-697)));
      }
    };
    static fm16(p0, globalState) {
      return function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(!(false)))).Elements) {
          let _10_v0 = _compr_7;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(!(false)))).contains(_10_v0)) {
            _coll7.push([_10_v0,(new BigNumber((_dafny.Set.fromElements(true, true)).length)).isEqualTo(new BigNumber(-608))]);
          }
        }
        return _coll7;
      }();
    };
    static fm17(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(true)).IsDisjointFrom(_dafny.MultiSet.fromElements(true, false)),new _dafny.CodePoint('d'.codePointAt(0)));
    };
    static m0(p0, globalState) {
      let r0 = [];
      let _hi0 = p0;
      for (let _11_i0 = new BigNumber((_dafny.Seq.UnicodeFromString("r")).length); _11_i0.isLessThan(_hi0); _11_i0 = _11_i0.plus(_dafny.ONE)) {
        let _12_v0;
        _12_v0 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_11_i0)).length), p0, p0);
        let _13_v1;
        _13_v1 = false;
        let _14_v2;
        let _nw0 = new _module.C0();
        _nw0.__ctor();
        _14_v2 = _nw0;
        let _15_v3;
        _15_v3 = _dafny.Map.Empty.slice().updateUnsafe(_14_v2,new BigNumber(803));
        let _16_v4;
        _16_v4 = _module.D3.create_DC7(_15_v3);
        let _17_v5;
        _17_v5 = _dafny.Map.Empty.slice().updateUnsafe(_13_v1,_16_v4);
        let _rhs0 = function () {
          let _coll8 = new _dafny.Set();
          for (const _compr_8 of ((_12_v0).Union(_12_v0)).Elements) {
            let _18_v6 = _compr_8;
            if (((_12_v0).Union(_12_v0)).contains(_18_v6)) {
              _coll8.add((_18_v6).plus((new BigNumber(201)).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-511)), function (_19_i1) {
                return new BigNumber(-128);
              }))).cardinality()))));
            }
          }
          return _coll8;
        }();
        let _rhs1 = _17_v5;
        _12_v0 = _rhs0;
        _17_v5 = _rhs1;
        (globalState).f11 = (_dafny.ZERO).minus(_11_i0);
        let _20_v7;
        let _init0 = ((_21_v1) => function (_22_i2) {
          return _21_v1;
        })(_13_v1);
        let _nw1 = Array((new BigNumber(21)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _20_v7 = _nw1;
        let _23_v8;
        _23_v8 = new _dafny.CodePoint('t'.codePointAt(0));
        let _24_v9;
        _24_v9 = _dafny.Map.Empty.slice().updateUnsafe(_23_v8,_20_v7);
        let _25_v10;
        _25_v10 = _module.D0.create_DC1(new BigNumber(443), new BigNumber(315), p0, p0, (((_24_v9).contains(_23_v8)) ? ((_24_v9).get(_23_v8)) : (_20_v7)));
        let _26_v11;
        _26_v11 = _dafny.Map.Empty.slice().updateUnsafe(_25_v10,_20_v7);
        let _27_v12;
        let _nw2 = Array((new BigNumber(17)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _20_v7;
        _nw2[(_dafny.ONE).toNumber()] = _20_v7;
        _nw2[(new BigNumber(2)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(3)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(4)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(5)).toNumber()] = (((_26_v11).contains(_25_v10)) ? ((_26_v11).get(_25_v10)) : (_20_v7));
        _nw2[(new BigNumber(6)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(7)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(8)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(9)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(10)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(11)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(12)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(13)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(14)).toNumber()] = _20_v7;
        _nw2[(new BigNumber(15)).toNumber()] = (_25_v10).dtor_cf5;
        _nw2[(new BigNumber(16)).toNumber()] = _20_v7;
        _27_v12 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_27_v12).length));
        (_27_v12)[_index0] = (_25_v10).dtor_cf5;
        let _index1 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_27_v12).length));
        (_27_v12)[_index1] = _20_v7;
        let _28_v13;
        let _nw3 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _28_v13 = _nw3;
        r0 = _28_v13;
      }
      let _29_v14;
      _29_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
      let _30_v15;
      _30_v15 = false;
      let _31_v16;
      _31_v16 = _dafny.Seq.of(false, !(_30_v15));
      let _32_v17;
      _32_v17 = _dafny.Set.fromElements(p0);
      let _33_v18;
      let _nw4 = Array((new BigNumber(16)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = _30_v15;
      _nw4[(_dafny.ONE).toNumber()] = false;
      _nw4[(new BigNumber(2)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(3)).toNumber()] = false;
      _nw4[(new BigNumber(4)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(5)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(6)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(7)).toNumber()] = (((_29_v14).contains(p0)) ? ((_29_v14).get(p0)) : (_30_v15));
      _nw4[(new BigNumber(8)).toNumber()] = !(_30_v15);
      _nw4[(new BigNumber(9)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(10)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(11)).toNumber()] = !(_30_v15);
      _nw4[(new BigNumber(12)).toNumber()] = _30_v15;
      _nw4[(new BigNumber(13)).toNumber()] = !(_30_v15);
      _nw4[(new BigNumber(14)).toNumber()] = !(_30_v15);
      _nw4[(new BigNumber(15)).toNumber()] = (_31_v16)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_32_v17).length)), new BigNumber((_31_v16).length))];
      _33_v18 = _nw4;
      let _34_v19;
      _34_v19 = _module.D0.create_DC1(_module.__default.fm0(globalState), new BigNumber((((false) ? (_29_v14) : (_29_v14))).length), p0, p0, _33_v18);
      let _source0 = _34_v19;
      if (_source0.is_DC1) {
        let _35___mcc_h0 = (_source0).cf1;
        let _36___mcc_h1 = (_source0).cf2;
        let _37___mcc_h2 = (_source0).cf3;
        let _38___mcc_h3 = (_source0).cf4;
        let _39___mcc_h4 = (_source0).cf5;
        let _40_cf5 = _39___mcc_h4;
        let _41_cf4 = _38___mcc_h3;
        let _42_cf3 = _37___mcc_h2;
        let _43_cf2 = _36___mcc_h1;
        let _44_cf1 = _35___mcc_h0;
        let _45_v20;
        _45_v20 = _dafny.Set.fromElements(true);
        let _46_v21;
        let _nw5 = new _module.C1();
        _nw5.__ctor((new BigNumber((_45_v20).length)).minus(_43_cf2));
        _46_v21 = _nw5;
        let _47_v22;
        let _init1 = ((_48_v15, _49_v16) => function (_50_i3) {
          return (_50_i3).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_48_v15,new BigNumber((_49_v16).length))).length));
        })(_30_v15, _31_v16);
        let _nw6 = Array((new BigNumber(26)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
          _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _47_v22 = _nw6;
        r0 = _47_v22;
        _41_cf4 = _46_v21.f18;
        (globalState).f12 = _30_v15;
      } else if (_source0.is_DC2) {
        let _51___mcc_h5 = (_source0).cf6;
        let _52_cf6 = _51___mcc_h5;
        let _53_v23;
        let _nw7 = new _module.C1();
        _nw7.__ctor(new BigNumber(-998));
        _53_v23 = _nw7;
        _53_v23 = _53_v23;
        _33_v18 = _33_v18;
        _52_cf6 = ((!(_30_v15) || (_30_v15)) ? (_52_cf6) : ((p0).isLessThanOrEqualTo(p0)));
        let _54_v24;
        let _nw8 = new _module.C0();
        _nw8.__ctor();
        _54_v24 = _nw8;
      } else if (_source0.is_DC3) {
        let _55___mcc_h6 = (_source0).cf7;
        let _56___mcc_h7 = (_source0).cf8;
        let _57___mcc_h8 = (_source0).cf9;
        let _58___mcc_h9 = (_source0).cf10;
        let _59_cf10 = _58___mcc_h9;
        let _60_cf9 = _57___mcc_h8;
        let _61_cf8 = _56___mcc_h7;
        let _62_cf7 = _55___mcc_h6;
        let _63_v25;
        _63_v25 = _dafny.Seq.UnicodeFromString("xbjalwtvy");
        let _64_v26;
        _64_v26 = _module.D0.create_DC0(new BigNumber((_63_v25).length));
        _64_v26 = _64_v26;
        (globalState).f12 = _dafny.Seq.IsPrefixOf(_63_v25, _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm8(globalState), _module.__default.safeIndex(_62_cf7, new BigNumber((_module.__default.fm8(globalState)).length)), new _dafny.CodePoint('l'.codePointAt(0))), _63_v25));
        let _65_v27;
        _65_v27 = _dafny.Set.fromElements(_30_v15, (_31_v16)[_module.__default.safeIndex(_62_cf7, new BigNumber((_31_v16).length))], _61_cf8, _30_v15, _30_v15);
        let _66_v28;
        _66_v28 = new _dafny.CodePoint('s'.codePointAt(0));
        _65_v27 = _module.__default.fm12(_66_v28, p0, globalState);
        (globalState).f12 = _61_cf8;
      } else {
        let _67___mcc_h10 = (_source0).cf0;
        let _68_cf0 = _67___mcc_h10;
        _30_v15 = _30_v15;
        let _69_v29;
        _69_v29 = _dafny.Seq.UnicodeFromString("dw");
        let _70_v30;
        _70_v30 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("igmkken"), _69_v29);
        if (((_31_v16)[_module.__default.safeIndex(new BigNumber(105), new BigNumber((_31_v16).length))]) === (_dafny.Seq.contains(_70_v30, _69_v29))) {
          let _71_v31;
          _71_v31 = _dafny.Seq.of(p0, _68_cf0);
          let _72_v32;
          _72_v32 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_71_v31);
          let _73_v33;
          _73_v33 = _module.D3.create_DC8(_71_v31, true, _30_v15, _69_v29, _30_v15);
          _72_v32 = (_72_v32).update(_module.__default.fm1(_30_v15, _30_v15, p0, globalState), (_73_v33).dtor_cf16);
          let _74_v34;
          _74_v34 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Concat(_69_v29, _69_v29)).length), ((_30_v15) ? (_68_cf0) : (p0)), (p0).multipliedBy(p0), p0, p0);
          let _75_v35;
          _75_v35 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_30_v15, _30_v15, p0, globalState),(_68_cf0).isLessThan(_module.__default.fm0(globalState)));
          let _76_v36;
          _76_v36 = _dafny.Set.fromElements(_30_v15);
          let _77_v37;
          _77_v37 = _dafny.MultiSet.fromElements(_31_v16);
          let _78_v38;
          _78_v38 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_77_v37);
          let _79_v39;
          _79_v39 = _dafny.Map.Empty.slice().updateUnsafe((_76_v36).Difference(_76_v36),((((_78_v38).contains(_30_v15)) ? ((_78_v38).get(_30_v15)) : (_77_v37))).IsDisjointFrom(_77_v37));
          let _80_v40;
          _80_v40 = new _dafny.CodePoint('b'.codePointAt(0));
          let _81_v41;
          _81_v41 = _dafny.Set.fromElements(_80_v40);
          let _82_v42;
          _82_v42 = _module.D0.create_DC2(_30_v15);
          let _rhs2 = (((_79_v39).contains(_76_v36)) ? ((_79_v39).get(_76_v36)) : (_30_v15));
          let _rhs3 = _74_v34;
          let _rhs4 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,!(_30_v15));
          let _rhs5 = (_81_v41).equals((_dafny.Set.fromElements(_80_v40, _80_v40, _80_v40, _80_v40, _80_v40)).Difference(_81_v41));
          let _rhs6 = ((((_82_v42).dtor_cf6) && (_30_v15)) ? (_68_cf0) : (_68_cf0));
          let _lhs0 = globalState;
          let _lhs1 = globalState;
          let _lhs2 = globalState;
          _lhs0.f12 = _rhs2;
          _74_v34 = _rhs3;
          _75_v35 = _rhs4;
          _lhs1.f12 = _rhs5;
          _lhs2.f11 = _rhs6;
          let _83_v43;
          let _nw9 = new _module.C0();
          _nw9.__ctor();
          _83_v43 = _nw9;
          let _84_v44;
          _84_v44 = _module.D4.create_DC9(_83_v43);
          let _85_v45;
          _85_v45 = _module.D4.create_DC11(_84_v44);
          let _86_v46;
          _86_v46 = _module.D4.create_DC11(_84_v44);
          let _87_v47;
          _87_v47 = _module.D4.create_DC11(_86_v46);
          let _rhs7 = new BigNumber((_69_v29).length);
          let _rhs8 = _87_v47;
          let _rhs9 = _80_v40;
          let _lhs3 = globalState;
          _lhs3.f11 = _rhs7;
          _87_v47 = _rhs8;
          _80_v40 = _rhs9;
          let _88_v48;
          let _nw10 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _88_v48 = _nw10;
          let _89_v49;
          _89_v49 = _dafny.Seq.of(_81_v41);
          let _index2 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_88_v48).length));
          (_88_v48)[_index2] = _89_v49;
          let _90_v50;
          _90_v50 = _dafny.MultiSet.fromElements(_30_v15);
          let _index3 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_88_v48).length));
          let _rhs10 = _33_v18;
          let _rhs11 = _89_v49;
          let _rhs12 = (_90_v50).Intersect(_90_v50);
          let _lhs4 = _88_v48;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_88_v48).length));
          _33_v18 = _rhs10;
          _lhs4[_lhs5] = _rhs11;
          _90_v50 = _rhs12;
          let _91_v51;
          let _init2 = ((_92_p0) => function (_93_i4) {
            return (_93_i4).plus(_92_p0);
          })(p0);
          let _nw11 = Array((new BigNumber(21)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
            _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _91_v51 = _nw11;
          let _index4 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_91_v51).length));
          (_91_v51)[_index4] = _module.__default.safeModuloInt(_68_cf0, _68_cf0);
          let _index5 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_91_v51).length));
          let _rhs13 = p0;
          let _rhs14 = new BigNumber((_31_v16).length);
          let _rhs15 = _68_cf0;
          let _rhs16 = (((_30_v15) ? (new BigNumber(-722)) : (_68_cf0))).plus(_module.__default.fm0(globalState));
          let _rhs17 = _module.__default.safeDivisionInt((_71_v31)[_module.__default.safeIndex(p0, new BigNumber((_71_v31).length))], (_71_v31)[_module.__default.safeIndex(new BigNumber((_69_v29).length), new BigNumber((_71_v31).length))]);
          let _lhs6 = globalState;
          let _lhs7 = globalState;
          let _lhs8 = _91_v51;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_91_v51).length));
          let _lhs10 = globalState;
          _lhs6.f11 = _rhs13;
          _68_cf0 = _rhs14;
          _lhs7.f11 = _rhs15;
          _lhs8[_lhs9] = _rhs16;
          _lhs10.f10 = _rhs17;
        } else {
          (globalState).f12 = (_module.D3.create_DC8(_dafny.Seq.of(p0), _module.__default.fm1(_30_v15, !(_30_v15), _68_cf0, globalState), _30_v15, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_94_i5) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}), _module.__default.safeIndex(_68_cf0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_95_i5) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length)), (_69_v29)[_module.__default.safeIndex(p0, new BigNumber((_69_v29).length))]), true)).dtor_cf20;
          (globalState).f12 = _30_v15;
          let _96_v52;
          _96_v52 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,(_68_cf0).isEqualTo(new BigNumber(491)));
          _96_v52 = (_96_v52).update(!_dafny.areEqual(_69_v29, _69_v29), _30_v15);
          _33_v18 = _33_v18;
          let _97_v53;
          let _init3 = ((_98_cf0, _99_p0) => function (_100_i6) {
            return (_100_i6).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_98_cf0,_99_p0)).length));
          })(_68_cf0, p0);
          let _nw12 = Array((new BigNumber(11)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw12.length); _i0_3++) {
            _nw12[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _97_v53 = _nw12;
          let _index6 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_97_v53).length));
          (_97_v53)[_index6] = p0;
          let _index7 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_97_v53).length));
          (_97_v53)[_index7] = p0;
        }
        _68_cf0 = (p0).minus(_68_cf0);
        let _101_v54;
        _101_v54 = _dafny.MultiSet.fromElements(_30_v15, false);
        (globalState).f12 = (new BigNumber((_101_v54).cardinality())).isEqualTo(_68_cf0);
      }
      if ((true) && (_30_v15)) {
        (globalState).f11 = p0;
        (globalState).f12 = !(_30_v15);
        let _102_v55;
        _102_v55 = new _dafny.CodePoint('v'.codePointAt(0));
        let _103_v56;
        _103_v56 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_102_v55);
        let _104_v57;
        _104_v57 = _dafny.Map.Empty.slice().updateUnsafe(_103_v56,_30_v15);
        let _index8 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length));
        (_33_v18)[_index8] = (((_104_v57).contains(_module.__default.fm17(_30_v15, globalState))) ? ((_104_v57).get(_module.__default.fm17(_30_v15, globalState))) : (_30_v15));
        let _105_v58;
        _105_v58 = _dafny.MultiSet.fromElements(p0);
        let _106_v59;
        _106_v59 = _dafny.Seq.of(p0);
        let _107_v60;
        _107_v60 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_30_v15);
        let _index9 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length));
        (_33_v18)[_index9] = ((((_105_v58).contains(p0)) ? ((_105_v58).get(p0)) : (p0))).isLessThanOrEqualTo((_106_v59)[_module.__default.safeIndex(new BigNumber((_107_v60).length), new BigNumber((_106_v59).length))]);
        let _108_v61;
        _108_v61 = _dafny.MultiSet.fromElements((_33_v18)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length))]);
        let _109_v62;
        _109_v62 = _dafny.Seq.UnicodeFromString("fjuprfrne");
        let _110_v63;
        _110_v63 = _dafny.Set.fromElements((_108_v61).update(true, _module.__default.abs(new BigNumber((_109_v62).length))), _108_v61, _108_v61);
        let _111_v64;
        _111_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_110_v63).length));
        let _112_v65;
        let _nw13 = new _module.C1();
        _nw13.__ctor(new BigNumber((_111_v64).length));
        _112_v65 = _nw13;
        let _113_v66;
        let _init4 = ((_114_v55) => function (_115_i7) {
          return _114_v55;
        })(_102_v55);
        let _nw14 = Array((new BigNumber(3)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw14.length); _i0_4++) {
          _nw14[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _113_v66 = _nw14;
        let _116_v67;
        _116_v67 = _module.D6.create_DC14(_30_v15, _102_v55, _113_v66);
        let _pat_let_tv0 = _102_v55;
        let _pat_let_tv1 = _30_v15;
        let _pat_let_tv2 = _30_v15;
        let _source1 = function (_pat_let0_0) {
          return function (_117_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_118_dt__update_hcf29_h0) {
                return function (_pat_let2_0) {
                  return function (_119_dt__update_hcf28_h0) {
                    return _module.D6.create_DC14(_119_dt__update_hcf28_h0, _118_dt__update_hcf29_h0, (_117_dt__update__tmp_h0).dtor_cf30);
                  }(_pat_let2_0);
                }((_pat_let_tv1) && (_pat_let_tv2));
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_116_v67);
        if (_source1.is_DC14) {
          let _120___mcc_h11 = (_source1).cf28;
          let _121___mcc_h12 = (_source1).cf29;
          let _122___mcc_h13 = (_source1).cf30;
          let _123_cf30 = _122___mcc_h13;
          let _124_cf29 = _121___mcc_h12;
          let _125_cf28 = _120___mcc_h11;
          (globalState).f10 = _112_v65.f18;
          _107_v60 = (_107_v60).update(!(_30_v15), (_33_v18)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length))]);
          let _126_v68;
          let _nw15 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _126_v68 = _nw15;
          let _index10 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_126_v68).length));
          (_126_v68)[_index10] = _109_v62;
          let _index11 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_126_v68).length));
          let _rhs18 = _102_v55;
          let _rhs19 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("jgxwsdkj"), _module.__default.safeIndex(_112_v65.f18, new BigNumber((_dafny.Seq.UnicodeFromString("jgxwsdkj")).length)), _124_cf29);
          let _lhs11 = _126_v68;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_126_v68).length));
          _124_cf29 = _rhs18;
          _lhs11[_lhs12] = _rhs19;
          let _127_v69;
          let _nw16 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _127_v69 = _nw16;
          let _index12 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_127_v69).length));
          (_127_v69)[_index12] = _112_v65.f18;
          let _index13 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_127_v69).length));
          (_127_v69)[_index13] = (_dafny.ZERO).minus(p0);
        } else {
          let _128___mcc_h14 = (_source1).cf27;
          let _129_cf27 = _128___mcc_h14;
          (globalState).f15 = _102_v55;
          let _130_v70;
          let _nw17 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _130_v70 = _nw17;
          let _131_v71;
          let _nw18 = Array((new BigNumber(28)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _130_v70;
          _nw18[(_dafny.ONE).toNumber()] = _130_v70;
          _nw18[(new BigNumber(2)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(3)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(4)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(5)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(6)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(7)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(8)).toNumber()] = ((_30_v15) ? (_130_v70) : (_130_v70));
          _nw18[(new BigNumber(9)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(10)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(11)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(12)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(13)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(14)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(15)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(16)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(17)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(18)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(19)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(20)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(21)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(22)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(23)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(24)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(25)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(26)).toNumber()] = _130_v70;
          _nw18[(new BigNumber(27)).toNumber()] = _130_v70;
          _131_v71 = _nw18;
          let _index14 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_131_v71).length));
          (_131_v71)[_index14] = _130_v70;
          let _132_v72;
          let _nw19 = new _module.C0();
          _nw19.__ctor();
          _132_v72 = _nw19;
          let _133_v73;
          let _nw20 = Array((new BigNumber(9)).toNumber()).fill([]);
          _133_v73 = _nw20;
          let _134_v74;
          let _init5 = ((_135_v14) => function (_136_i8) {
            return _135_v14;
          })(_29_v14);
          let _nw21 = Array((new BigNumber(5)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw21.length); _i0_5++) {
            _nw21[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _134_v74 = _nw21;
          let _index15 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_133_v73).length));
          (_133_v73)[_index15] = _134_v74;
          let _137_v75;
          _137_v75 = _dafny.Set.fromElements(_130_v70, _130_v70);
          let _index16 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_131_v71).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_133_v73).length));
          let _rhs20 = _130_v70;
          let _rhs21 = _132_v72;
          let _rhs22 = (((_137_v75).IsProperSubsetOf(_137_v75)) ? (_134_v74) : (_134_v74));
          let _lhs13 = _131_v71;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_131_v71).length));
          let _lhs15 = _133_v73;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_133_v73).length));
          _lhs13[_lhs14] = _rhs20;
          _132_v72 = _rhs21;
          _lhs15[_lhs16] = _rhs22;
          let _138_v76;
          let _nw22 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _138_v76 = _nw22;
          let _139_v77;
          _139_v77 = _dafny.Map.Empty.slice().updateUnsafe(_109_v62,_112_v65.f18);
          let _index18 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_138_v76).length));
          (_138_v76)[_index18] = _139_v77;
          let _index19 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_138_v76).length));
          (_138_v76)[_index19] = _139_v77;
          let _140_v78;
          _140_v78 = _dafny.MultiSet.fromElements(_109_v62);
          let _141_v79;
          _141_v79 = _dafny.Map.Empty.slice().updateUnsafe((_33_v18)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length))],p0);
          let _index20 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_130_v70).length));
          (_130_v70)[_index20] = (((_140_v78).contains((_132_v72).fm7(_30_v15, new BigNumber((_dafny.Seq.of(_102_v55, _102_v55)).length), _141_v79, _109_v62, globalState))) ? ((_140_v78).get((_132_v72).fm7(_30_v15, new BigNumber((_dafny.Seq.of(_102_v55, _102_v55)).length), _141_v79, _109_v62, globalState))) : ((_dafny.ZERO).minus(new BigNumber((_108_v61).cardinality()))));
          let _142_v81;
          _142_v81 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll9 = new _dafny.Set();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-954), new BigNumber(990))) {
              let _143_v80 = _compr_9;
              if (((new BigNumber(-954)).isLessThanOrEqualTo(_143_v80)) && ((_143_v80).isLessThan(new BigNumber(990)))) {
                _coll9.add((_143_v80).multipliedBy(_112_v65.f18));
              }
            }
            return _coll9;
          }(),p0);
          let _144_v83;
          _144_v83 = _module.D6.create_DC13(_129_cf27);
          let _145_v84;
          _145_v84 = _dafny.MultiSet.fromElements(_144_v83);
          let _146_v85;
          _146_v85 = _dafny.Map.Empty.slice().updateUnsafe(_145_v84,(_33_v18)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length))]);
          let _index21 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length));
          let _index22 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_130_v70).length));
          let _rhs23 = _106_v59;
          let _rhs24 = !(_30_v15);
          let _rhs25 = (((_142_v81).contains(function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_29_v14).Keys.Elements) {
              let _148_v82 = _compr_11;
              if ((_29_v14).contains(_148_v82)) {
                _coll11.add(_module.__default.safeModuloInt(_148_v82, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(420))).length)));
              }
            }
            return _coll11;
          }())) ? ((_142_v81).get(function () {
            let _coll10 = new _dafny.Set();
            for (const _compr_10 of (_29_v14).Keys.Elements) {
              let _147_v82 = _compr_10;
              if ((_29_v14).contains(_147_v82)) {
                _coll10.add(_module.__default.safeModuloInt(_147_v82, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(420))).length)));
              }
            }
            return _coll10;
          }())) : (p0));
          let _rhs26 = (_112_v65.f18).minus(_module.__default.safeModuloInt(new BigNumber((_146_v85).length), new BigNumber((_dafny.Set.fromElements(_30_v15)).length)));
          let _rhs27 = _dafny.Seq.Concat(_109_v62, _109_v62);
          let _lhs17 = _33_v18;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_33_v18).length));
          let _lhs19 = globalState;
          let _lhs20 = _130_v70;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_130_v70).length));
          _106_v59 = _rhs23;
          _lhs17[_lhs18] = _rhs24;
          _lhs19.f10 = _rhs25;
          _lhs20[_lhs21] = _rhs26;
          _109_v62 = _rhs27;
        }
      } else {
        let _149_v86;
        _149_v86 = _dafny.MultiSet.fromElements(!(_30_v15));
        let _150_v87;
        _150_v87 = _dafny.Map.Empty.slice().updateUnsafe(_33_v18,_33_v18);
        let _rhs28 = _30_v15;
        let _rhs29 = _149_v86;
        let _rhs30 = new BigNumber(((_32_v17).Union(_32_v17)).length);
        let _rhs31 = (((p0).isLessThan(p0)) ? ((_150_v87).Merge(_150_v87)) : (_150_v87));
        let _lhs22 = globalState;
        let _lhs23 = globalState;
        _lhs22.f12 = _rhs28;
        _149_v86 = _rhs29;
        _lhs23.f11 = _rhs30;
        _150_v87 = _rhs31;
        let _151_v88;
        _151_v88 = _dafny.Set.fromElements(true);
        (globalState).f10 = (p0).plus(new BigNumber((_151_v88).length));
        let _152_v89;
        _152_v89 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_30_v15);
        let _153_v90;
        _153_v90 = _dafny.MultiSet.fromElements(_152_v89, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_30_v15, _30_v15, p0, globalState),_30_v15));
        _153_v90 = (((_31_v16)[_module.__default.safeIndex(p0, new BigNumber((_31_v16).length))]) ? ((_153_v90).update(_152_v89, _module.__default.abs(p0))) : ((_153_v90).Union(_153_v90)));
        let _154_v91;
        let _nw23 = Array((new BigNumber(19)).toNumber()).fill([]);
        _154_v91 = _nw23;
        let _155_v92;
        let _nw24 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _155_v92 = _nw24;
        let _index23 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_154_v91).length));
        (_154_v91)[_index23] = _155_v92;
        let _156_v93;
        _156_v93 = _dafny.Seq.UnicodeFromString("hpyho");
        let _157_v94;
        _157_v94 = _156_v93;
        let _158_v95;
        _158_v95 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index24 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_154_v91).length));
        let _nw25 = Array((new BigNumber(5)).toNumber());
        _nw25[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lksj"), _156_v93), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lksj"), _156_v93)).length)), (_156_v93)[_module.__default.safeIndex(p0, new BigNumber((_156_v93).length))]);
        _nw25[(_dafny.ONE).toNumber()] = (_157_v94);
        _nw25[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_156_v93, _156_v93);
        _nw25[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_159_i9) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        });
        _nw25[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_156_v93, _156_v93), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_156_v93, _156_v93)).length)), _158_v95);
        (_154_v91)[_index24] = _nw25;
        let _rhs32 = _30_v15;
        let _rhs33 = _158_v95;
        let _rhs34 = p0;
        let _rhs35 = _dafny.Seq.Concat(_module.__default.fm8(globalState), _dafny.Seq.update(_156_v93, _module.__default.safeIndex(p0, new BigNumber((_156_v93).length)), _158_v95));
        let _lhs24 = globalState;
        let _lhs25 = globalState;
        let _lhs26 = globalState;
        _lhs24.f12 = _rhs32;
        _158_v95 = _rhs33;
        _lhs25.f10 = _rhs34;
        _lhs26.f3 = _rhs35;
      }
      let _160_i10;
      _160_i10 = _dafny.ZERO;
      L0: {
        while (!(_30_v15)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_160_i10)) {
              break L0;
            }
            _160_i10 = (_160_i10).plus(_dafny.ONE);
            let _161_v96;
            let _nw26 = new _module.C0();
            _nw26.__ctor();
            _161_v96 = _nw26;
            let _162_v97;
            _162_v97 = _dafny.Map.Empty.slice().updateUnsafe(_161_v96,_30_v15);
            let _163_v98;
            _163_v98 = _dafny.Map.Empty.slice().updateUnsafe(_30_v15,_162_v97);
            _163_v98 = (_163_v98).update(_30_v15, (_162_v97).Merge(_162_v97));
            (globalState).f11 = (new BigNumber((_module.__default.fm15(p0, _30_v15, new BigNumber(794), globalState)).length)).minus(p0);
            let _164_v99;
            _164_v99 = _dafny.Seq.of(p0, p0);
            let _165_v100;
            let _nw27 = new _module.C1();
            _nw27.__ctor(new BigNumber(451));
            _165_v100 = _nw27;
            let _166_v102;
            _166_v102 = new _dafny.CodePoint('u'.codePointAt(0));
            let _167_v103;
            _167_v103 = _dafny.Seq.of(_166_v102, _166_v102);
            let _168_v104;
            _168_v104 = _dafny.Map.Empty.slice().updateUnsafe(_165_v100,function () {
              let _coll12 = new _dafny.Map();
              for (const _compr_12 of (_167_v103).Elements) {
                let _169_v101 = _compr_12;
                if (_dafny.Seq.contains(_167_v103, _169_v101)) {
                  _coll12.push([_169_v101,_165_v100.f18]);
                }
              }
              return _coll12;
            }());
            let _170_v105;
            let _nw28 = Array((new BigNumber(23)).toNumber());
            _nw28[(_dafny.ZERO).toNumber()] = ((_30_v15) ? (p0) : (new BigNumber(899)));
            _nw28[(_dafny.ONE).toNumber()] = (p0).multipliedBy(p0);
            _nw28[(new BigNumber(2)).toNumber()] = new BigNumber(586);
            _nw28[(new BigNumber(3)).toNumber()] = new BigNumber(120);
            _nw28[(new BigNumber(4)).toNumber()] = p0;
            _nw28[(new BigNumber(5)).toNumber()] = p0;
            _nw28[(new BigNumber(6)).toNumber()] = p0;
            _nw28[(new BigNumber(7)).toNumber()] = p0;
            _nw28[(new BigNumber(8)).toNumber()] = p0;
            _nw28[(new BigNumber(9)).toNumber()] = p0;
            _nw28[(new BigNumber(10)).toNumber()] = _module.__default.fm0(globalState);
            _nw28[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update(_164_v99, _module.__default.safeIndex(p0, new BigNumber((_164_v99).length)), (_dafny.ZERO).minus(p0))).length);
            _nw28[(new BigNumber(12)).toNumber()] = p0;
            _nw28[(new BigNumber(13)).toNumber()] = p0;
            _nw28[(new BigNumber(14)).toNumber()] = (new BigNumber((_dafny.Seq.update(_164_v99, _module.__default.safeIndex(p0, new BigNumber((_164_v99).length)), new BigNumber(472))).length)).multipliedBy(new BigNumber((_168_v104).length));
            _nw28[(new BigNumber(15)).toNumber()] = _165_v100.f18;
            _nw28[(new BigNumber(16)).toNumber()] = _165_v100.f18;
            _nw28[(new BigNumber(17)).toNumber()] = p0;
            _nw28[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_167_v103, _module.__default.safeIndex(p0, new BigNumber((_167_v103).length)), (_167_v103)[_module.__default.safeIndex(p0, new BigNumber((_167_v103).length))])).length);
            _nw28[(new BigNumber(19)).toNumber()] = _165_v100.f18;
            _nw28[(new BigNumber(20)).toNumber()] = ((_module.__default.fm1(_30_v15, _30_v15, new BigNumber((_167_v103).length), globalState)) ? (_module.__default.fm0(globalState)) : ((_164_v99)[_module.__default.safeIndex(_165_v100.f18, new BigNumber((_164_v99).length))]));
            _nw28[(new BigNumber(21)).toNumber()] = _165_v100.f18;
            _nw28[(new BigNumber(22)).toNumber()] = p0;
            _170_v105 = _nw28;
            r0 = _170_v105;
            (globalState).f12 = !(_30_v15);
          }
        }
      }
      _33_v18 = _33_v18;
      let _171_v106;
      let _nw29 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _171_v106 = _nw29;
      _171_v106 = _171_v106;
      let _172_v107;
      let _nw30 = Array((_dafny.ONE).toNumber());
      _nw30[(_dafny.ZERO).toNumber()] = p0;
      _172_v107 = _nw30;
      r0 = _172_v107;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _173_v0;
      _173_v0 = _dafny.Seq.UnicodeFromString("yo");
      let _174_v2;
      _174_v2 = new _dafny.CodePoint('k'.codePointAt(0));
      let _175_v3;
      _175_v3 = _dafny.Set.fromElements(_174_v2);
      let _176_v4;
      _176_v4 = _dafny.MultiSet.fromElements(function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)))).Elements) {
          let _177_v1 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0))), _177_v1)) {
            _coll13.add(_177_v1);
          }
        }
        return _coll13;
      }(), _175_v3);
      let _178_globalState;
      let _nw31 = new _module.GlobalState();
      _nw31.__ctor(new BigNumber(134), true, new BigNumber(107), _173_v0, _dafny.Seq.UnicodeFromString("ci"), new BigNumber(380), new BigNumber(-147), false, false, new BigNumber(727), new BigNumber(540), new BigNumber(576), false, false, _176_v4, new _dafny.CodePoint('c'.codePointAt(0)), new BigNumber(403), new BigNumber(-68));
      _178_globalState = _nw31;
      let _179_v5;
      _179_v5 = new BigNumber(980);
      let _hi1 = _179_v5;
      for (let _180_i0 = new BigNumber(-735); _180_i0.isLessThan(_hi1); _180_i0 = _180_i0.plus(_dafny.ONE)) {
        let _181_v6;
        let _nw32 = Array((new BigNumber(7)).toNumber()).fill(false);
        _181_v6 = _nw32;
        let _182_v7;
        _182_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_181_v6, _181_v6, _181_v6),false);
        let _183_v8;
        _183_v8 = _dafny.Seq.of(_181_v6, _181_v6, _181_v6, _181_v6, _181_v6);
        _182_v7 = (_182_v7).update(_dafny.Seq.Concat(_183_v8, _183_v8), false);
        let _184_v9;
        let _nw33 = Array((new BigNumber(2)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = new BigNumber(-13);
        _nw33[(_dafny.ONE).toNumber()] = (_module.D0.create_DC0(_180_i0)).dtor_cf0;
        _184_v9 = _nw33;
        let _185_v10;
        _185_v10 = true;
        let _186_v11;
        _186_v11 = _dafny.Seq.of(!(_185_v10), _185_v10);
        let _187_v12;
        _187_v12 = _dafny.Map.Empty.slice().updateUnsafe((_179_v5).multipliedBy(_180_i0),_185_v10);
        let _rhs36 = (((new BigNumber(880)).isLessThan(new BigNumber(-438))) ? (_184_v9) : (_184_v9));
        let _rhs37 = !(!_dafny.areEqual(_173_v0, _173_v0));
        let _rhs38 = ((_185_v10) ? (_185_v10) : ((_186_v11)[_module.__default.safeIndex(_179_v5, new BigNumber((_186_v11).length))]));
        let _rhs39 = (((_187_v12).contains(new BigNumber(716))) ? ((_187_v12).get(new BigNumber(716))) : (_185_v10));
        let _rhs40 = ((_dafny.ZERO).minus(_180_i0)).minus((new BigNumber((_173_v0).length)).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(_178_globalState)))));
        let _lhs27 = _178_globalState;
        let _lhs28 = _178_globalState;
        let _lhs29 = _178_globalState;
        _184_v9 = _rhs36;
        _lhs27.f12 = _rhs37;
        _lhs28.f12 = _rhs38;
        _lhs29.f12 = _rhs39;
        _179_v5 = _rhs40;
        let _188_v13;
        _188_v13 = _dafny.MultiSet.fromElements(_module.__default.fm0(_178_globalState));
        _188_v13 = (_188_v13).Union(_188_v13);
        let _189_v14;
        _189_v14 = _dafny.MultiSet.fromElements(_185_v10, true, true);
        let _190_v15;
        _190_v15 = _dafny.Map.Empty.slice().updateUnsafe(_185_v10,new BigNumber((_189_v14).cardinality()));
        let _191_v16;
        _191_v16 = _dafny.Map.Empty.slice().updateUnsafe(_179_v5,_190_v15);
        _191_v16 = (_191_v16).update(_180_i0, _190_v15);
      }
      let _192_v17;
      _192_v17 = false;
      let _193_v18;
      _193_v18 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_179_v5);
      let _194_v19;
      _194_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_179_v5, _179_v5, new BigNumber((_193_v18).length))).length),!(_192_v17));
      if ((((_194_v19).contains(_179_v5)) ? ((_194_v19).get(_179_v5)) : ((_179_v5).isEqualTo(_179_v5)))) {
        let _195_v20;
        let _nw34 = Array((new BigNumber(22)).toNumber()).fill(false);
        _195_v20 = _nw34;
        let _196_v21;
        _196_v21 = _module.D0.create_DC3(_179_v5, _192_v17, _179_v5, new BigNumber((_173_v0).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length));
        (_195_v20)[_index25] = (((((_194_v19).contains(_179_v5)) ? ((_194_v19).get(_179_v5)) : ((_196_v21).dtor_cf8))) ? (_module.__default.fm1(_192_v17, _192_v17, _179_v5, _178_globalState)) : (_192_v17));
        let _197_v22;
        _197_v22 = _dafny.Set.fromElements(_192_v17, !(_192_v17));
        let _index26 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length));
        (_195_v20)[_index26] = _module.__default.fm1(false, (_197_v22).IsDisjointFrom(_197_v22), _module.__default.safeModuloInt(_179_v5, _179_v5), _178_globalState);
        let _198_v23;
        let _nw35 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
        _198_v23 = _nw35;
        let _199_v25;
        _199_v25 = _dafny.Map.Empty.slice().updateUnsafe(_179_v5,_179_v5);
        let _200_v26;
        _200_v26 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_199_v25).Keys.Elements) {
            let _201_v24 = _compr_14;
            if ((_199_v25).contains(_201_v24)) {
              _coll14.push([(_201_v24).multipliedBy(_179_v5),_module.D0.create_DC2((_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))])]);
            }
          }
          return _coll14;
        }(),_179_v5);
        let _index27 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_198_v23).length));
        (_198_v23)[_index27] = _200_v26;
        let _index28 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_198_v23).length));
        (_198_v23)[_index28] = _module.__default.fm2(_178_globalState);
        let _source2 = _module.__default.fm3(new BigNumber((_173_v0).length), _178_globalState);
        if (_source2.is_DC1) {
          let _202___mcc_h0 = (_source2).cf1;
          let _203___mcc_h1 = (_source2).cf2;
          let _204___mcc_h2 = (_source2).cf3;
          let _205___mcc_h3 = (_source2).cf4;
          let _206___mcc_h4 = (_source2).cf5;
          let _207_cf5 = _206___mcc_h4;
          let _208_cf4 = _205___mcc_h3;
          let _209_cf3 = _204___mcc_h2;
          let _210_cf2 = _203___mcc_h1;
          let _211_cf1 = _202___mcc_h0;
          let _212_v27;
          let _out0;
          _out0 = _module.__default.m0((((_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))]) ? (_210_cf2) : (_208_cf4)), _178_globalState);
          _212_v27 = _out0;
          (_178_globalState).f10 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_208_cf4));
          let _213_v28;
          let _nw36 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
          _213_v28 = _nw36;
          let _214_v29;
          _214_v29 = _dafny.Seq.of(_210_cf2);
          let _215_v30;
          _215_v30 = _dafny.Set.fromElements(_214_v29);
          let _index29 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_213_v28).length));
          (_213_v28)[_index29] = _215_v30;
          let _index30 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_213_v28).length));
          (_213_v28)[_index30] = _215_v30;
          let _index31 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length));
          (_195_v20)[_index31] = !(_dafny.Seq.contains(_173_v0, _174_v2));
        } else if (_source2.is_DC2) {
          let _216___mcc_h5 = (_source2).cf6;
          let _217_cf6 = _216___mcc_h5;
          let _218_v31;
          _218_v31 = _dafny.Seq.of((_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))]);
          _218_v31 = _dafny.Seq.Concat(_218_v31, _218_v31);
          let _rhs41 = _module.__default.safeModuloInt(_179_v5, _179_v5);
          let _rhs42 = ((_dafny.ZERO).minus(_179_v5)).minus(_179_v5);
          let _lhs30 = _178_globalState;
          let _lhs31 = _178_globalState;
          _lhs30.f10 = _rhs41;
          _lhs31.f10 = _rhs42;
          (_178_globalState).f12 = false;
          let _219_v32;
          let _init6 = ((_220_v5) => function (_221_i1) {
            return (_221_i1).plus(_220_v5);
          })(_179_v5);
          let _nw37 = Array((new BigNumber(4)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw37.length); _i0_6++) {
            _nw37[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _219_v32 = _nw37;
          let _index32 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_219_v32).length));
          (_219_v32)[_index32] = _179_v5;
          let _222_v33;
          _222_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_217_cf6),(_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))]);
          let _index33 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_219_v32).length));
          let _rhs43 = ((_217_cf6) ? (_179_v5) : (_module.__default.safeModuloInt(_179_v5, new BigNumber((_222_v33).length))));
          let _rhs44 = _dafny.Seq.Concat(_173_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(401)), ((_223_v2) => function (_224_i2) {
            return _223_v2;
          })(_174_v2)));
          let _lhs32 = _219_v32;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_219_v32).length));
          _lhs32[_lhs33] = _rhs43;
          _173_v0 = _rhs44;
        } else if (_source2.is_DC3) {
          let _225___mcc_h6 = (_source2).cf7;
          let _226___mcc_h7 = (_source2).cf8;
          let _227___mcc_h8 = (_source2).cf9;
          let _228___mcc_h9 = (_source2).cf10;
          let _229_cf10 = _228___mcc_h9;
          let _230_cf9 = _227___mcc_h8;
          let _231_cf8 = _226___mcc_h7;
          let _232_cf7 = _225___mcc_h6;
          let _233_v34;
          _233_v34 = _dafny.Seq.of(_194_v19, (_194_v19).update(new BigNumber(959), _231_cf8), _194_v19);
          let _234_v35;
          _234_v35 = _dafny.MultiSet.fromElements(_174_v2);
          let _235_v36;
          _235_v36 = _dafny.Seq.of(_231_cf8);
          let _236_v37;
          _236_v37 = _dafny.MultiSet.fromElements(_179_v5, _229_cf10);
          let _237_v38;
          let _nw38 = Array((new BigNumber(29)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_179_v5, _232_cf7);
          _nw38[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(((_231_cf8) ? (_179_v5) : (_229_cf10)));
          _nw38[(new BigNumber(2)).toNumber()] = new BigNumber(11);
          _nw38[(new BigNumber(3)).toNumber()] = (_232_cf7).plus(_230_cf9);
          _nw38[(new BigNumber(4)).toNumber()] = _230_cf9;
          _nw38[(new BigNumber(5)).toNumber()] = (_module.__default.fm0(_178_globalState)).multipliedBy(new BigNumber((_233_v34).length));
          _nw38[(new BigNumber(6)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(7)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), ((_238_cf9) => function (_239_i3) {
            return _238_cf9;
          })(_230_cf9))).length);
          _nw38[(new BigNumber(9)).toNumber()] = (_230_cf9).multipliedBy(new BigNumber(419));
          _nw38[(new BigNumber(10)).toNumber()] = new BigNumber(437);
          _nw38[(new BigNumber(11)).toNumber()] = (new BigNumber(((_234_v35).update(_174_v2, _module.__default.abs(_179_v5))).cardinality())).minus(_232_cf7);
          _nw38[(new BigNumber(12)).toNumber()] = _232_cf7;
          _nw38[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm4(_178_globalState), _235_v36)).length);
          _nw38[(new BigNumber(14)).toNumber()] = _232_cf7;
          _nw38[(new BigNumber(15)).toNumber()] = new BigNumber(-540);
          _nw38[(new BigNumber(16)).toNumber()] = _229_cf10;
          _nw38[(new BigNumber(17)).toNumber()] = (((_236_v37).contains(_229_cf10)) ? ((_236_v37).get(_229_cf10)) : (_229_cf10));
          _nw38[(new BigNumber(18)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(19)).toNumber()] = _229_cf10;
          _nw38[(new BigNumber(20)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(21)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_173_v0).length))).minus(new BigNumber((_module.__default.fm5(_229_cf10, _178_globalState)).length)));
          _nw38[(new BigNumber(23)).toNumber()] = _229_cf10;
          _nw38[(new BigNumber(24)).toNumber()] = (_179_v5).multipliedBy(new BigNumber(-530));
          _nw38[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Seq.of(_179_v5, (_dafny.ZERO).minus(new BigNumber((_173_v0).length)))).length);
          _nw38[(new BigNumber(26)).toNumber()] = _229_cf10;
          _nw38[(new BigNumber(27)).toNumber()] = _179_v5;
          _nw38[(new BigNumber(28)).toNumber()] = _229_cf10;
          _237_v38 = _nw38;
          let _index34 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_237_v38).length));
          (_237_v38)[_index34] = (_230_cf9).plus(_232_cf7);
          let _index35 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_237_v38).length));
          (_237_v38)[_index35] = new BigNumber((_173_v0).length);
          let _240_v39;
          _240_v39 = _module.D1.create_DC4(_174_v2);
          let _pat_let_tv3 = _174_v2;
          (_178_globalState).f15 = (function (_pat_let3_0) {
            return function (_241_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_242_dt__update_hcf11_h0) {
                  return _module.D1.create_DC4(_242_dt__update_hcf11_h0);
                }(_pat_let4_0);
              }(_pat_let_tv3);
            }(_pat_let3_0);
          }(_240_v39)).dtor_cf11;
          let _243_v40;
          let _nw39 = Array((new BigNumber(25)).toNumber()).fill(_module.D0.Default());
          _243_v40 = _nw39;
          let _244_v41;
          _244_v41 = _dafny.Set.fromElements(_243_v40);
          let _index36 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_237_v38).length));
          (_237_v38)[_index36] = new BigNumber(((_244_v41).Union(_244_v41)).length);
          _231_cf8 = ((_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))]) && (_192_v17);
        } else {
          let _245___mcc_h10 = (_source2).cf0;
          let _246_cf0 = _245___mcc_h10;
          (_178_globalState).f12 = !(_192_v17) || ((_195_v20)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length))]);
          (_178_globalState).f3 = _173_v0;
          (_178_globalState).f15 = _174_v2;
          let _247_v42;
          let _out1;
          _out1 = _module.__default.m0(_module.__default.fm0(_178_globalState), _178_globalState);
          _247_v42 = _out1;
        }
        let _248_v43;
        _248_v43 = _dafny.MultiSet.fromElements(new BigNumber((_173_v0).length));
        let _index37 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_195_v20).length));
        (_195_v20)[_index37] = !((_248_v43).IsSubsetOf((_248_v43).Union(_248_v43)));
        let _249_v44;
        _249_v44 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,!(_192_v17));
        let _250_v45;
        let _out2;
        _out2 = _module.__default.m0((((_248_v43).contains(new BigNumber((_249_v44).length))) ? ((_248_v43).get(new BigNumber((_249_v44).length))) : (_179_v5)), _178_globalState);
        _250_v45 = _out2;
      } else {
        let _251_v46;
        let _out3;
        _out3 = _module.__default.m0(_179_v5, _178_globalState);
        _251_v46 = _out3;
        let _252_v47;
        _252_v47 = _module.D0.create_DC3(_179_v5, _192_v17, new BigNumber(-428), _179_v5);
        if ((_252_v47).dtor_cf8) {
          let _253_v48;
          _253_v48 = _173_v0;
          (_178_globalState).f3 = (_253_v48);
          let _index38 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_251_v46).length));
          (_251_v46)[_index38] = _179_v5;
          let _index39 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_251_v46).length));
          let _rhs45 = _179_v5;
          let _rhs46 = _192_v17;
          let _rhs47 = (new BigNumber((_173_v0).length)).minus(_179_v5);
          let _rhs48 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_173_v0, _173_v0), _173_v0);
          let _lhs34 = _251_v46;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_251_v46).length));
          let _lhs36 = _178_globalState;
          _179_v5 = _rhs45;
          _192_v17 = _rhs46;
          _lhs34[_lhs35] = _rhs47;
          _lhs36.f12 = _rhs48;
          let _254_v49;
          let _out4;
          _out4 = _module.__default.m0((_251_v46)[_module.__default.safeIndex(new BigNumber(421), new BigNumber((_251_v46).length))], _178_globalState);
          _254_v49 = _out4;
          let _255_v50;
          _255_v50 = _module.D0.create_DC2(_192_v17);
          let _pat_let_tv4 = _192_v17;
          let _256_v51;
          _256_v51 = _dafny.Seq.of(function (_pat_let5_0) {
            return function (_257_dt__update__tmp_h1) {
              return function (_pat_let6_0) {
                return function (_258_dt__update_hcf6_h0) {
                  return _module.D0.create_DC2(_258_dt__update_hcf6_h0);
                }(_pat_let6_0);
              }(!(_pat_let_tv4));
            }(_pat_let5_0);
          }(_255_v50), _255_v50);
          _256_v51 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), ((_259_v50) => function (_260_i4) {
            return _259_v50;
          })(_255_v50));
          let _261_v52;
          _261_v52 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_192_v17);
          let _262_v53;
          _262_v53 = _dafny.Set.fromElements((((_261_v52).contains(_192_v17)) ? ((_261_v52).get(_192_v17)) : (_192_v17)));
          let _263_v54;
          let _out5;
          _out5 = _module.__default.m0(((!(!(_192_v17))) ? (new BigNumber(97)) : (new BigNumber((_262_v53).length))), _178_globalState);
          _263_v54 = _out5;
        } else {
          let _rhs49 = (_179_v5).plus(_179_v5);
          let _rhs50 = _dafny.Seq.UnicodeFromString("hfnyb");
          let _lhs37 = _178_globalState;
          let _lhs38 = _178_globalState;
          _lhs37.f11 = _rhs49;
          _lhs38.f3 = _rhs50;
          let _264_v55;
          _264_v55 = _dafny.Set.fromElements(_192_v17, _192_v17);
          let _265_v56;
          _265_v56 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_264_v55);
          (_178_globalState).f10 = (_179_v5).minus(new BigNumber((_265_v56).length));
          let _266_v57;
          let _out6;
          _out6 = _module.__default.m0(_179_v5, _178_globalState);
          _266_v57 = _out6;
          let _267_v58;
          let _out7;
          _out7 = _module.__default.m0(_179_v5, _178_globalState);
          _267_v58 = _out7;
          (_178_globalState).f12 = _192_v17;
        }
        let _268_v59;
        _268_v59 = _dafny.Seq.of(_192_v17, true, _192_v17);
        let _269_v60;
        let _nw40 = Array((new BigNumber(8)).toNumber());
        _nw40[(_dafny.ZERO).toNumber()] = _192_v17;
        _nw40[(_dafny.ONE).toNumber()] = (_179_v5).isLessThan(_179_v5);
        _nw40[(new BigNumber(2)).toNumber()] = _192_v17;
        _nw40[(new BigNumber(3)).toNumber()] = _192_v17;
        _nw40[(new BigNumber(4)).toNumber()] = _192_v17;
        _nw40[(new BigNumber(5)).toNumber()] = !((((_194_v19).contains(_179_v5)) ? ((_194_v19).get(_179_v5)) : (_192_v17)));
        _nw40[(new BigNumber(6)).toNumber()] = !((_268_v59)[_module.__default.safeIndex(_179_v5, new BigNumber((_268_v59).length))]);
        _nw40[(new BigNumber(7)).toNumber()] = _192_v17;
        _269_v60 = _nw40;
        let _index40 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length));
        (_269_v60)[_index40] = _192_v17;
        let _270_v61;
        _270_v61 = _dafny.MultiSet.fromElements(_192_v17);
        let _index41 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length));
        (_269_v60)[_index41] = _module.__default.fm1(_192_v17, !(_dafny.MultiSet.fromElements(new BigNumber((_270_v61).cardinality()), _179_v5, _179_v5)).equals((_dafny.MultiSet.fromElements(_179_v5, _179_v5)).update(_179_v5, _module.__default.abs(_179_v5))), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("wbxk")).length), _179_v5)), _178_globalState);
        let _271_v62;
        let _out8;
        _out8 = _module.__default.m0(new BigNumber(5), _178_globalState);
        _271_v62 = _out8;
        let _index42 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length));
        (_269_v60)[_index42] = !(!((_269_v60)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length))])) || (!((_269_v60)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length))]) || ((_269_v60)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_269_v60).length))]));
      }
      let _272_v63;
      let _nw41 = Array((new BigNumber(7)).toNumber()).fill(false);
      _272_v63 = _nw41;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_272_v63).length))) {
        let _273_i5 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_273_i5)) && ((_273_i5).isLessThan(new BigNumber((_272_v63).length))))) {
          (_272_v63)[(_273_i5)] = !(_dafny.MultiSet.fromElements(_179_v5)).contains(new BigNumber(((_dafny.MultiSet.fromElements(_173_v0, _173_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), ((_274_v2) => function (_275_i6) {
            return _274_v2;
          })(_174_v2)))).Intersect(_dafny.MultiSet.fromElements(_173_v0))).cardinality()));
        }
      }
      let _276_v64;
      _276_v64 = _dafny.Set.fromElements(_192_v17, _module.__default.fm1(_192_v17, _192_v17, new BigNumber((_dafny.Seq.UnicodeFromString("yhaahdov")).length), _178_globalState), _192_v17);
      _192_v17 = (_276_v64).IsProperSubsetOf(_276_v64);
      let _277_v65;
      _277_v65 = _dafny.Map.Empty.slice().updateUnsafe(_272_v63,_179_v5);
      let _278_v66;
      _278_v66 = _dafny.Seq.of(!(_module.__default.fm1(!(!(_192_v17)), _192_v17, new BigNumber((_277_v65).length), _178_globalState)), _192_v17, _192_v17, _192_v17, _192_v17);
      let _279_v67;
      _279_v67 = _dafny.Set.fromElements(_179_v5, new BigNumber(706), _179_v5);
      let _280_v68;
      _280_v68 = _module.D0.create_DC3(_179_v5, (new BigNumber((_278_v66).length)).isLessThan(_179_v5), _179_v5, (new BigNumber((_279_v67).length)).multipliedBy(_179_v5));
      _280_v68 = _280_v68;
      let _281_v69;
      _281_v69 = _dafny.Map.Empty.slice().updateUnsafe(_193_v18,_192_v17);
      let _hi2 = new BigNumber(66);
      for (let _282_i7 = (_dafny.ZERO).minus(new BigNumber((_281_v69).length)); _282_i7.isLessThan(_hi2); _282_i7 = _282_i7.plus(_dafny.ONE)) {
        let _283_v70;
        let _nw42 = new _module.C1();
        _nw42.__ctor(_282_i7);
        _283_v70 = _nw42;
        (_178_globalState).f12 = (!(_192_v17) || (_192_v17)) === (_192_v17);
        (_283_v70).m2(_272_v63, _dafny.Seq.Concat(_173_v0, _173_v0), new BigNumber((_dafny.Seq.of(_283_v70.f18)).length), _178_globalState);
        if (!(_192_v17) || (_192_v17)) {
          let _nw43 = new _module.C1();
          _nw43.__ctor(_283_v70.f18);
          _283_v70 = _nw43;
          let _284_v71;
          _284_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(806),_283_v70);
          let _285_v72;
          _285_v72 = _dafny.Seq.of(_284_v71);
          let _286_v73;
          _286_v73 = _dafny.Seq.of((_285_v72)[_module.__default.safeIndex(_283_v70.f18, new BigNumber((_285_v72).length))]);
          (_178_globalState).f12 = ((_286_v73)[_module.__default.safeIndex(_282_i7, new BigNumber((_286_v73).length))]).equals(_284_v71);
          let _287_v74;
          let _nw44 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _287_v74 = _nw44;
          let _288_v75;
          _288_v75 = _dafny.Map.Empty.slice().updateUnsafe(false,_192_v17);
          let _289_v76;
          _289_v76 = _dafny.Map.Empty.slice().updateUnsafe(_282_i7,new BigNumber(556));
          let _290_v77;
          let _nw45 = Array((new BigNumber(22)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _179_v5;
          _nw45[(_dafny.ONE).toNumber()] = new BigNumber((_288_v75).length);
          _nw45[(new BigNumber(2)).toNumber()] = _282_i7;
          _nw45[(new BigNumber(3)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(4)).toNumber()] = _282_i7;
          _nw45[(new BigNumber(5)).toNumber()] = _179_v5;
          _nw45[(new BigNumber(6)).toNumber()] = _282_i7;
          _nw45[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-670));
          _nw45[(new BigNumber(8)).toNumber()] = _179_v5;
          _nw45[(new BigNumber(9)).toNumber()] = new BigNumber(297);
          _nw45[(new BigNumber(10)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(11)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(12)).toNumber()] = _282_i7;
          _nw45[(new BigNumber(13)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(14)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(15)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(16)).toNumber()] = _283_v70.f18;
          _nw45[(new BigNumber(17)).toNumber()] = _179_v5;
          _nw45[(new BigNumber(18)).toNumber()] = _module.__default.fm0(_178_globalState);
          _nw45[(new BigNumber(19)).toNumber()] = new BigNumber((_289_v76).length);
          _nw45[(new BigNumber(20)).toNumber()] = _282_i7;
          _nw45[(new BigNumber(21)).toNumber()] = _179_v5;
          _290_v77 = _nw45;
          let _291_v78;
          _291_v78 = _dafny.Set.fromElements(_290_v77);
          let _292_v79;
          _292_v79 = _dafny.Map.Empty.slice().updateUnsafe((_291_v78).Union(_291_v78),(new BigNumber((_193_v18).length)).multipliedBy(_179_v5));
          let _rhs51 = (((_292_v79).contains(_291_v78)) ? ((_292_v79).get(_291_v78)) : (new BigNumber(110)));
          let _rhs52 = _287_v74;
          let _rhs53 = _module.__default.safeDivisionInt(_283_v70.f18, _179_v5);
          let _rhs54 = _192_v17;
          let _lhs39 = _178_globalState;
          let _lhs40 = _178_globalState;
          _lhs39.f10 = _rhs51;
          _287_v74 = _rhs52;
          _179_v5 = _rhs53;
          _lhs40.f12 = _rhs54;
          let _293_v80;
          _293_v80 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_283_v70);
          let _294_v81;
          let _nw46 = new _module.C1();
          _nw46.__ctor(new BigNumber(((_293_v80).Merge(_293_v80)).length));
          _294_v81 = _nw46;
          (_178_globalState).f11 = _283_v70.f18;
        } else {
          let _295_v82;
          let _nw47 = new _module.C0();
          _nw47.__ctor();
          _295_v82 = _nw47;
          let _296_v83;
          _296_v83 = _dafny.Map.Empty.slice().updateUnsafe(_295_v82,_283_v70.f18);
          let _297_v84;
          _297_v84 = _module.D3.create_DC7(_296_v83);
          let _298_v85;
          _298_v85 = _dafny.Map.Empty.slice().updateUnsafe(_297_v84,_282_i7);
          (_283_v70).f18 = _module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_297_v84,_283_v70.f18)).Merge(_298_v85)).length), new BigNumber((_173_v0).length));
          let _299_v86;
          _299_v86 = _dafny.MultiSet.fromElements(_295_v82, _295_v82);
          let _300_v87;
          let _nw48 = new _module.C1();
          _nw48.__ctor(new BigNumber((_299_v86).cardinality()));
          _300_v87 = _nw48;
          let _index43 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_272_v63).length));
          (_272_v63)[_index43] = _192_v17;
          let _index44 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_272_v63).length));
          (_272_v63)[_index44] = (_300_v87.f18).isLessThanOrEqualTo(_283_v70.f18);
          let _301_v88;
          let _nw49 = new _module.C0();
          _nw49.__ctor();
          _301_v88 = _nw49;
          (_178_globalState).f10 = (_dafny.ZERO).minus(_283_v70.f18);
        }
      }
      let _302_v89;
      _302_v89 = _dafny.Seq.of(_179_v5, new BigNumber(-221), _179_v5, _179_v5, new BigNumber((_173_v0).length));
      let _hi3 = _179_v5;
      for (let _303_i8 = new BigNumber((_302_v89).length); _303_i8.isLessThan(_hi3); _303_i8 = _303_i8.plus(_dafny.ONE)) {
        let _304_v90;
        let _nw50 = new _module.C0();
        _nw50.__ctor();
        _304_v90 = _nw50;
        _192_v17 = _192_v17;
        let _305_v91;
        _305_v91 = _dafny.Map.Empty.slice().updateUnsafe(true,_173_v0);
        let _306_v92;
        let _nw51 = new _module.C1();
        _nw51.__ctor(_module.__default.safeDivisionInt(new BigNumber(741), new BigNumber((_305_v91).length)));
        _306_v92 = _nw51;
        let _307_v93;
        _307_v93 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_272_v63);
        _307_v93 = _307_v93;
      }
      let _308_v94;
      let _init7 = ((_309_v0, _310_v2) => function (_311_i9) {
        return _dafny.Seq.Concat(_309_v0, _dafny.Seq.update(_309_v0, _module.__default.safeIndex(new BigNumber((_309_v0).length), new BigNumber((_309_v0).length)), _310_v2));
      })(_173_v0, _174_v2);
      let _nw52 = Array((new BigNumber(7)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw52.length); _i0_7++) {
        _nw52[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _308_v94 = _nw52;
      _308_v94 = _308_v94;
      (_178_globalState).f12 = false;
      let _312_v95;
      let _nw53 = Array((new BigNumber(23)).toNumber()).fill([]);
      _312_v95 = _nw53;
      _312_v95 = ((_192_v17) ? (((_192_v17) ? (_312_v95) : (_312_v95))) : (_312_v95));
      (_178_globalState).f10 = _179_v5;
      let _313_v96;
      _313_v96 = _dafny.Map.Empty.slice().updateUnsafe(_179_v5,_179_v5);
      (_178_globalState).f12 = !(_179_v5).isEqualTo((new BigNumber(-135)).plus(new BigNumber((_dafny.Seq.update(_module.__default.fm4(_178_globalState), _module.__default.safeIndex(new BigNumber((_313_v96).length), new BigNumber((_module.__default.fm4(_178_globalState)).length)), _192_v17)).length)));
      let _index45 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_308_v94).length));
      (_308_v94)[_index45] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), ((_314_v2) => function (_315_i10) {
        return _314_v2;
      })(_174_v2));
      let _316_v97;
      _316_v97 = _dafny.Seq.of(_173_v0);
      let _index46 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_308_v94).length));
      (_308_v94)[_index46] = (_316_v97)[_module.__default.safeIndex(_179_v5, new BigNumber((_316_v97).length))];
      if (((_192_v17) ? (_192_v17) : (_192_v17))) {
        let _317_v98;
        let _out9;
        _out9 = _module.__default.m0(_179_v5, _178_globalState);
        _317_v98 = _out9;
        (_178_globalState).f12 = (_179_v5).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_179_v5, _179_v5));
        (_178_globalState).f10 = (((_313_v96).contains(new BigNumber(465))) ? ((_313_v96).get(new BigNumber(465))) : (new BigNumber((_302_v89).length)));
        (_178_globalState).f10 = _179_v5;
        let _318_v99;
        let _nw54 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _318_v99 = _nw54;
        let _index47 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_318_v99).length));
        (_318_v99)[_index47] = _dafny.Seq.of(_179_v5, _module.__default.fm0(_178_globalState));
        let _index48 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_318_v99).length));
        (_318_v99)[_index48] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-759)), _dafny.Seq.of(_179_v5, _179_v5, new BigNumber(877)));
      } else {
        _179_v5 = (_179_v5).minus(new BigNumber((_278_v66).length));
        let _319_v100;
        _319_v100 = _dafny.Seq.of(_272_v63, _272_v63, _272_v63);
        let _320_v101;
        let _nw55 = new _module.C1();
        _nw55.__ctor(new BigNumber((((_192_v17) ? (_319_v100) : (_319_v100))).length));
        _320_v101 = _nw55;
        _320_v101 = _320_v101;
        _320_v101 = _320_v101;
        if (!((_179_v5).isLessThanOrEqualTo(_179_v5))) {
          let _321_v102;
          let _nw56 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _321_v102 = _nw56;
          let _322_v103;
          let _nw57 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Set.Empty);
          _322_v103 = _nw57;
          let _323_v104;
          _323_v104 = _322_v103;
          let _324_v105;
          _324_v105 = _dafny.Set.fromElements(_323_v104, _323_v104, _322_v103, _323_v104);
          let _index49 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_321_v102).length));
          (_321_v102)[_index49] = ((_192_v17) ? (_324_v105) : (_324_v105));
          let _index50 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_321_v102).length));
          (_321_v102)[_index50] = _324_v105;
          let _325_v106;
          let _out10;
          _out10 = _module.__default.m0(_179_v5, _178_globalState);
          _325_v106 = _out10;
          _192_v17 = _192_v17;
          let _index51 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_325_v106).length));
          (_325_v106)[_index51] = _320_v101.f18;
          let _index52 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_325_v106).length));
          let _rhs55 = _320_v101.f18;
          let _rhs56 = _272_v63;
          let _lhs41 = _325_v106;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_325_v106).length));
          _lhs41[_lhs42] = _rhs55;
          _272_v63 = _rhs56;
          let _326_v107;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _326_v107 = _nw58;
          let _327_v108;
          _327_v108 = _dafny.Map.Empty.slice().updateUnsafe(!(_192_v17),_326_v107);
          (_178_globalState).f12 = _module.__default.fm1(!((new BigNumber(((_327_v108).update(_192_v17, _326_v107)).length)).isEqualTo(new BigNumber((_279_v67).length))), _192_v17, _179_v5, _178_globalState);
        } else {
          (_320_v101).m2(_272_v63, _173_v0, _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), ((_328_v2) => function (_329_i11) {
            return _328_v2;
          })(_174_v2)))).cardinality()), new BigNumber(243)), _178_globalState);
          _278_v66 = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm1(_192_v17, _192_v17, _179_v5, _178_globalState)), _dafny.Seq.Concat(_278_v66, _278_v66));
          (_178_globalState).f11 = (_dafny.ZERO).minus(new BigNumber(((_308_v94)[_module.__default.safeIndex(new BigNumber(551), new BigNumber((_308_v94).length))]).length));
          (_178_globalState).f12 = ((_dafny.ZERO).minus(_179_v5)).isEqualTo((new BigNumber((_193_v18).length)).minus(_179_v5));
          let _330_v109;
          let _nw59 = new _module.C1();
          _nw59.__ctor(_320_v101.f18);
          _330_v109 = _nw59;
        }
        let _331_v110;
        _331_v110 = _dafny.Seq.of(_320_v101, _320_v101, _320_v101, _320_v101);
        _320_v101 = (_331_v110)[_module.__default.safeIndex(_179_v5, new BigNumber((_331_v110).length))];
      }
      let _332_v111;
      _332_v111 = _dafny.MultiSet.fromElements(_192_v17, _192_v17, _module.__default.fm1(_192_v17, _192_v17, _179_v5, _178_globalState), _192_v17);
      let _333_v112;
      _333_v112 = _dafny.Map.Empty.slice().updateUnsafe(_192_v17,_192_v17);
      let _334_v113;
      _334_v113 = _dafny.Map.Empty.slice().updateUnsafe(_278_v66,!((((_333_v112).contains(_192_v17)) ? ((_333_v112).get(_192_v17)) : (_192_v17))));
      (_178_globalState).f12 = !(_module.__default.fm16(new BigNumber((_332_v111).cardinality()), _178_globalState)).equals(_334_v113);
      let _335_v114;
      let _nw60 = Array((new BigNumber(3)).toNumber());
      _335_v114 = _nw60;
      let _336_v115;
      let _nw61 = new _module.C0();
      _nw61.__ctor();
      _336_v115 = _nw61;
      let _index53 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_335_v114).length));
      (_335_v114)[_index53] = _336_v115;
      let _index54 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_335_v114).length));
      (_335_v114)[_index54] = _336_v115;
      process.stdout.write((_173_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_174_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v3).equals(_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v4).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_178_globalState.f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_178_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_globalState).f14).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_179_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_192_v17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-2161)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_194_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v63)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v64).equals(_dafny.Set.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_277_v65).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_278_v66, _dafny.Seq.of(true, false, false, false, false, false, false, false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v67).equals(_dafny.Set.fromElements(new BigNumber(-2161), new BigNumber(706)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v68).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v68).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v68).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v68).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v69).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-2161)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_302_v89, _dafny.Seq.of(_dafny.ONE, new BigNumber(-221), _dafny.ONE, _dafny.ONE, new BigNumber(403)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_308_v94)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_313_v96).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_316_v97, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yokkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_332_v111).equals(_dafny.MultiSet.fromElements(false, false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_333_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v113).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false, false, false, false, false, false, false, false, false, false),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf6) {
      let $dt = new D0(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf7, cf8, cf9, cf10) {
      let $dt = new D0(2);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5(cf12, cf13) {
      let $dt = new D1(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC4(cf11) {
      let $dt = new D1(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf14) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + this.cf14.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC8(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC7(cf15) {
      let $dt = new D3(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + this.cf19.toVerbatimString(true) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17 && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.Seq.of(), false, false, _dafny.Seq.UnicodeFromString(""), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf22, cf23, cf24) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC9(cf21) {
      let $dt = new D4(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf25) {
      let $dt = new D4(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(new _dafny.CodePoint('D'.codePointAt(0)), [], _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf26) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf28, cf29, cf30) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC13(cf27) {
      let $dt = new D6(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC14(false, new _dafny.CodePoint('D'.codePointAt(0)), []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf31) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf33, cf34) {
      let $dt = new D8(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC18(cf35, cf36, cf37) {
      let $dt = new D8(1);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC16(cf32) {
      let $dt = new D8(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf38) {
      let $dt = new D8(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC18" + "(" + this.cf35.toVerbatimString(true) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC17(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = _dafny.Seq.UnicodeFromString("");
      this.f10 = _dafny.ZERO;
      this.f11 = _dafny.ZERO;
      this.f12 = false;
      this.f15 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f0 = _dafny.ZERO;
      this._f1 = false;
      this._f2 = _dafny.ZERO;
      this._f4 = _dafny.Seq.UnicodeFromString("");
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = false;
      this._f9 = _dafny.ZERO;
      this._f13 = false;
      this._f14 = _dafny.MultiSet.Empty;
      this._f16 = _dafny.ZERO;
      this._f17 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements((_dafny.ZERO).minus((new BigNumber(-495)).multipliedBy(new BigNumber((_dafny.Set.fromElements(!(false), false)).length))), ((false) ? (new BigNumber(-643)) : (new BigNumber(993))));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cdmixbs"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("blliiliu"), _dafny.Seq.UnicodeFromString("vthendu")));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18) {
      let _this = this;
      (_this).f18 = f18;
      return;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _337_v0;
      let _nw62 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _337_v0 = _nw62;
      let _338_v1;
      _338_v1 = new _dafny.CodePoint('a'.codePointAt(0));
      let _index55 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_337_v0).length));
      (_337_v0)[_index55] = _338_v1;
      let _index56 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_337_v0).length));
      (_337_v0)[_index56] = new _dafny.CodePoint('n'.codePointAt(0));
      let _339_v2;
      let _nw63 = Array((new BigNumber(27)).toNumber()).fill(false);
      _339_v2 = _nw63;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_339_v2).length))) {
        let _340_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_340_i0)) && ((_340_i0).isLessThan(new BigNumber((_339_v2).length))))) {
          (_339_v2)[(_340_i0)] = p1;
        }
      }
      (_this).f18 = p0;
      let _341_v3;
      let _out11;
      _out11 = _module.__default.m0(p0, globalState);
      _341_v3 = _out11;
      let _hi4 = _module.__default.fm0(globalState);
      for (let _342_i1 = _this.f18; _342_i1.isLessThan(_hi4); _342_i1 = _342_i1.plus(_dafny.ONE)) {
        let _index57 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_339_v2).length));
        (_339_v2)[_index57] = true;
        let _343_v4;
        _343_v4 = _dafny.MultiSet.fromElements(_342_i1);
        let _index58 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_339_v2).length));
        (_339_v2)[_index58] = (_343_v4).IsProperSubsetOf(_343_v4);
        let _344_v5;
        let _nw64 = new _module.C0();
        _nw64.__ctor();
        _344_v5 = _nw64;
        let _345_v6;
        _345_v6 = _dafny.Seq.UnicodeFromString("bfkwvkpem");
        let _346_v7;
        _346_v7 = _dafny.Map.Empty.slice().updateUnsafe(_345_v6,new BigNumber(938));
        _346_v7 = _346_v7;
        let _index59 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_341_v3).length));
        (_341_v3)[_index59] = (_dafny.ZERO).minus(p3);
        let _index60 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_341_v3).length));
        (_341_v3)[_index60] = (_dafny.ZERO).minus(((_dafny.ZERO).minus((_this.f18).plus(_342_i1))).plus(_this.f18));
      }
      if (_module.__default.fm1(p1, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("rrme"), _module.__default.fm8(globalState)), p3, globalState)) {
        let _347_v8;
        _347_v8 = _dafny.Map.Empty.slice().updateUnsafe(_341_v3,_module.D0.create_DC0(p3));
        _347_v8 = _347_v8;
        let _348_v9;
        _348_v9 = _dafny.MultiSet.fromElements(true, p1);
        let _349_v10;
        _349_v10 = _dafny.Map.Empty.slice().updateUnsafe(_337_v0,(_348_v9).IsProperSubsetOf(_dafny.MultiSet.fromElements(p1, p1, p1, p1)));
        let _350_v11;
        _350_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_337_v0);
        _349_v10 = (_349_v10).update((((_350_v11).contains(p0)) ? ((_350_v11).get(p0)) : (_337_v0)), p1);
        (globalState).f11 = p0;
        (globalState).f12 = p1;
        let _351_v12;
        _351_v12 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-927)), function (_352_i2) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})).length));
        let _353_v13;
        _353_v13 = _dafny.Set.fromElements(_module.D0.create_DC0(p3), _351_v12, _module.D0.create_DC0(new BigNumber(537)));
        let _index61 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_339_v2).length));
        (_339_v2)[_index61] = (_353_v13).IsProperSubsetOf(_353_v13);
        let _index62 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_339_v2).length));
        let _rhs57 = false;
        let _rhs58 = p1;
        let _lhs43 = globalState;
        let _lhs44 = _339_v2;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_339_v2).length));
        _lhs43.f12 = _rhs57;
        _lhs44[_lhs45] = _rhs58;
      } else {
        let _354_v14;
        _354_v14 = _dafny.Seq.of((_337_v0)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_337_v0).length))], (_337_v0)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_337_v0).length))]);
        let _355_v15;
        _355_v15 = _354_v14;
        let _356_v16;
        _356_v16 = _dafny.Map.Empty.slice().updateUnsafe(_355_v15,p1);
        _356_v16 = _dafny.Map.Empty.slice().updateUnsafe(_354_v14,true);
        let _357_v17;
        _357_v17 = _dafny.Seq.of(p1);
        let _358_v18;
        _358_v18 = _dafny.Map.Empty.slice().updateUnsafe(_354_v14,_357_v17);
        _358_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_354_v14, _354_v14),_357_v17);
        let _359_v19;
        _359_v19 = _dafny.Map.Empty.slice().updateUnsafe(_339_v2,p1);
        let _360_v20;
        _360_v20 = _module.D1.create_DC4(_338_v1);
        let _rhs59 = (_359_v19).Merge(_359_v19);
        let _rhs60 = (_dafny.ZERO).minus(p0);
        let _rhs61 = p3;
        let _rhs62 = _360_v20;
        let _lhs46 = globalState;
        let _lhs47 = globalState;
        _359_v19 = _rhs59;
        _lhs46.f11 = _rhs60;
        _lhs47.f10 = _rhs61;
        _360_v20 = _rhs62;
        r0 = _339_v2;
        let _361_v21;
        _361_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _362_v22;
        _362_v22 = _dafny.Set.fromElements(_361_v21, _dafny.Map.Empty.slice().updateUnsafe(p1,p1), _361_v21);
        (globalState).f12 = (_362_v22).IsDisjointFrom(_362_v22);
      }
      r0 = _339_v2;
      r1 = p3;
      return [r0, r1];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _363_v0;
      let _nw65 = Array((new BigNumber(2)).toNumber());
      _nw65[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw65[(_dafny.ONE).toNumber()] = (_this.f18).minus(new BigNumber(-962));
      _363_v0 = _nw65;
      let _364_v1;
      _364_v1 = false;
      _363_v0 = ((_364_v1) ? (_363_v0) : (_363_v0));
      let _365_i0;
      _365_i0 = _dafny.ZERO;
      L1: {
        while (((true) ? ((new BigNumber(181)).isLessThanOrEqualTo(_this.f18)) : ((p2).isLessThanOrEqualTo(p2)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_365_i0)) {
              break L1;
            }
            _365_i0 = (_365_i0).plus(_dafny.ONE);
            let _366_v2;
            let _init8 = ((_367_v1, _368_p2) => function (_369_i1) {
              return ((_367_v1) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_370_v1) => function (_371_i2) {
                return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_370_v1, _370_v1, _370_v1))).cardinality());
              })(_367_v1))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(154)), ((_372_p2) => function (_373_i3) {
                return _372_p2;
              })(_368_p2))));
            })(_364_v1, p2);
            let _nw66 = Array((new BigNumber(25)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw66.length); _i0_8++) {
              _nw66[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _366_v2 = _nw66;
            let _374_v3;
            _374_v3 = _dafny.Seq.of(p2);
            let _index63 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_366_v2).length));
            (_366_v2)[_index63] = _374_v3;
            let _index64 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_366_v2).length));
            (_366_v2)[_index64] = _374_v3;
            (globalState).f12 = (((new BigNumber((p1).length)).isLessThanOrEqualTo(new BigNumber((p1).length))) ? (!(_364_v1)) : (_364_v1));
            let _375_v4;
            let _nw67 = new _module.C0();
            _nw67.__ctor();
            _375_v4 = _nw67;
            let _376_v5;
            _376_v5 = _dafny.Map.Empty.slice().updateUnsafe(_375_v4,(_this.f18).minus(p2));
            let _index65 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((p0).length));
            (p0)[_index65] = !(_364_v1);
            let _377_v6;
            _377_v6 = _module.D3.create_DC7(_376_v5);
            let _378_v7;
            _378_v7 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), function (_379_i4) {
  return _this.f18;
})).length), _this.f18, p2, _this.f18, p0);
            let _380_v8;
            _380_v8 = _dafny.Seq.of(_module.D0.create_DC1(new BigNumber(250), new BigNumber((_dafny.Seq.UnicodeFromString("mocct")).length), (_dafny.ZERO).minus(_this.f18), _this.f18, p0));
            let _index66 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((p0).length));
            let _rhs63 = (_377_v6).dtor_cf15;
            let _rhs64 = _364_v1;
            let _rhs65 = !_dafny.Seq.contains(_380_v8, _378_v7);
            let _rhs66 = new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_381_i5) {
              return new _dafny.CodePoint('x'.codePointAt(0));
            }))).length);
            let _lhs48 = p0;
            let _lhs49 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((p0).length));
            let _lhs50 = globalState;
            _376_v5 = _rhs63;
            _lhs48[_lhs49] = _rhs64;
            _364_v1 = _rhs65;
            _lhs50.f10 = _rhs66;
            let _382_v9;
            _382_v9 = _module.D0.create_DC2(_364_v1);
            let _383_v10;
            _383_v10 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((p0).length))],_382_v9);
            let _384_v11;
            _384_v11 = _dafny.Map.Empty.slice().updateUnsafe(_383_v10,_this.f18);
            let _385_v12;
            _385_v12 = new _dafny.CodePoint('c'.codePointAt(0));
            let _386_v13;
            _386_v13 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(p2, new BigNumber((p1).length)), _385_v12)).length));
            let _387_v15;
            _387_v15 = _dafny.Set.fromElements(_383_v10, _383_v10);
            (globalState).f12 = (_384_v11).equals((_dafny.Map.Empty.slice().updateUnsafe(_383_v10,new BigNumber((_386_v13).cardinality()))).Merge(function () {
              let _coll15 = new _dafny.Map();
              for (const _compr_15 of (_387_v15).Elements) {
                let _388_v14 = _compr_15;
                if ((_387_v15).contains(_388_v14)) {
                  _coll15.push([_388_v14,new BigNumber(2)]);
                }
              }
              return _coll15;
            }()));
          }
        }
      }
      let _389_v16;
      let _init9 = ((_390_p1, _391_p2, _392_v1) => function (_393_i7) {
        return _module.D3.create_DC8(_dafny.Seq.of(new BigNumber(896), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.update(_390_p1, _module.__default.safeIndex(new BigNumber(505), new BigNumber((_390_p1).length)), new _dafny.CodePoint('v'.codePointAt(0)))).length), _391_p2))).cardinality())), _392_v1, _392_v1, _390_p1, _392_v1);
      })(p1, p2, _364_v1);
      let _nw68 = Array((new BigNumber(17)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw68.length); _i0_9++) {
        _nw68[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _389_v16 = _nw68;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_389_v16).length))) {
        let _394_i6 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_394_i6)) && ((_394_i6).isLessThan(new BigNumber((_389_v16).length))))) {
          (_389_v16)[(_394_i6)] = ((false) ? (_module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), ((_395_v1) => function (_396_i8) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_395_v1,_this.f18)).length);
})(_364_v1)), _364_v1, _364_v1, p1, _364_v1)) : (_module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_397_p2) => function (_398_i9) {
  return _397_p2;
})(p2)), !(_364_v1), _364_v1, p1, _364_v1)));
        }
      }
      if (false) {
        if ((p2).isLessThanOrEqualTo(p2)) {
          let _399_v17;
          let _nw69 = new _module.C0();
          _nw69.__ctor();
          _399_v17 = _nw69;
          let _rhs67 = _364_v1;
          let _rhs68 = true;
          let _rhs69 = p2;
          let _rhs70 = p2;
          let _lhs51 = globalState;
          let _lhs52 = globalState;
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          _lhs51.f12 = _rhs67;
          _lhs52.f12 = _rhs68;
          _lhs53.f10 = _rhs69;
          _lhs54.f11 = _rhs70;
          let _400_v18;
          _400_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(_364_v1),p1);
          let _401_v19;
          _401_v19 = p1;
          let _402_v20;
          _402_v20 = _dafny.Seq.of(_this.f18);
          let _403_v21;
          _403_v21 = _module.D3.create_DC8(_402_v20, _364_v1, _364_v1, p1, _364_v1);
          let _404_v22;
          let _nw70 = Array((new BigNumber(29)).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("lglnyxquc");
          _nw70[(_dafny.ONE).toNumber()] = p1;
          _nw70[(new BigNumber(2)).toNumber()] = (((_400_v18).contains(_364_v1)) ? ((_400_v18).get(_364_v1)) : (p1));
          _nw70[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-156)), function (_405_i10) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          });
          _nw70[(new BigNumber(4)).toNumber()] = p1;
          _nw70[(new BigNumber(5)).toNumber()] = p1;
          _nw70[(new BigNumber(6)).toNumber()] = (p1);
          _nw70[(new BigNumber(7)).toNumber()] = p1;
          _nw70[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(738)), function (_406_i11) {
            return (_module.D1.create_DC4(new _dafny.CodePoint('l'.codePointAt(0)))).dtor_cf11;
          });
          _nw70[(new BigNumber(9)).toNumber()] = p1;
          _nw70[(new BigNumber(10)).toNumber()] = p1;
          _nw70[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(p1, (_403_v21).dtor_cf19);
          _nw70[(new BigNumber(12)).toNumber()] = p1;
          _nw70[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("hfqwixbp");
          _nw70[(new BigNumber(14)).toNumber()] = p1;
          _nw70[(new BigNumber(15)).toNumber()] = p1;
          _nw70[(new BigNumber(16)).toNumber()] = p1;
          _nw70[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(p1, p1);
          _nw70[(new BigNumber(18)).toNumber()] = p1;
          _nw70[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("kyqwj");
          _nw70[(new BigNumber(20)).toNumber()] = (_403_v21).dtor_cf19;
          _nw70[(new BigNumber(21)).toNumber()] = p1;
          _nw70[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("aqevck");
          _nw70[(new BigNumber(23)).toNumber()] = p1;
          _nw70[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(p1, p1);
          _nw70[(new BigNumber(25)).toNumber()] = p1;
          _nw70[(new BigNumber(26)).toNumber()] = p1;
          _nw70[(new BigNumber(27)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_407_i12) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          });
          _nw70[(new BigNumber(28)).toNumber()] = _dafny.Seq.UnicodeFromString("mgyrwllh");
          _404_v22 = _nw70;
          let _index67 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_404_v22).length));
          (_404_v22)[_index67] = p1;
          let _index68 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_404_v22).length));
          (_404_v22)[_index68] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), function (_408_i13) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          });
          let _409_v23;
          _409_v23 = _dafny.Seq.of(_399_v17);
          _399_v17 = (_409_v23)[_module.__default.safeIndex((_this.f18).minus((_dafny.ZERO).minus(_this.f18)), new BigNumber((_409_v23).length))];
          let _410_v24;
          let _nw71 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _410_v24 = _nw71;
          let _411_v25;
          _411_v25 = _dafny.Seq.of(_363_v0, _363_v0, _363_v0, _363_v0);
          let _index69 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_410_v24).length));
          (_410_v24)[_index69] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_363_v0, _363_v0, _363_v0), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_363_v0, _363_v0, _363_v0)).length)), _363_v0), _411_v25);
          let _index70 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_410_v24).length));
          (_410_v24)[_index70] = _dafny.Seq.Concat(_411_v25, _dafny.Seq.of(_363_v0, _363_v0));
        } else {
          _363_v0 = _363_v0;
          (globalState).f15 = _module.__default.fm9(_364_v1, _364_v1, globalState);
          let _412_v26;
          let _init10 = function (_413_i14) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          };
          let _nw72 = Array((new BigNumber(20)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw72.length); _i0_10++) {
            _nw72[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _412_v26 = _nw72;
          _412_v26 = _412_v26;
          _364_v1 = ((_364_v1) ? (_364_v1) : (_364_v1));
          let _414_v27;
          let _nw73 = new _module.C0();
          _nw73.__ctor();
          _414_v27 = _nw73;
        }
        let _415_v28;
        _415_v28 = new _dafny.CodePoint('h'.codePointAt(0));
        let _416_v29;
        _416_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(350)), function (_417_i15) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }),_415_v28);
        _416_v29 = (_416_v29).update(_dafny.Seq.Concat(p1, p1), _415_v28);
        let _418_v30;
        let _nw74 = new _module.C0();
        _nw74.__ctor();
        _418_v30 = _nw74;
        let _419_v31;
        _419_v31 = _dafny.Map.Empty.slice().updateUnsafe(_418_v30,(_dafny.ZERO).minus(p2));
        _419_v31 = (_419_v31).update(_418_v30, _this.f18);
        let _420_v32;
        _420_v32 = _module.D0.create_DC1(new BigNumber((p1).length), p2, p2, (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_this.f18), p2)), p0);
        let _source3 = _420_v32;
        if (_source3.is_DC1) {
          let _421___mcc_h0 = (_source3).cf1;
          let _422___mcc_h1 = (_source3).cf2;
          let _423___mcc_h2 = (_source3).cf3;
          let _424___mcc_h3 = (_source3).cf4;
          let _425___mcc_h4 = (_source3).cf5;
          let _426_cf5 = _425___mcc_h4;
          let _427_cf4 = _424___mcc_h3;
          let _428_cf3 = _423___mcc_h2;
          let _429_cf2 = _422___mcc_h1;
          let _430_cf1 = _421___mcc_h0;
          (globalState).f12 = (_module.__default.safeModuloInt(_427_cf4, _this.f18)).isLessThanOrEqualTo(_428_cf3);
          let _index71 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_426_cf5).length));
          (_426_cf5)[_index71] = _364_v1;
          let _431_v33;
          _431_v33 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-308)), ((_432_cf1) => function (_433_i16) {
            return _432_cf1;
          })(_430_cf1))));
          let _434_v34;
          _434_v34 = _dafny.MultiSet.fromElements(new BigNumber(853), _this.f18);
          let _435_v35;
          _435_v35 = _dafny.Seq.of(_364_v1, _364_v1);
          let _436_v36;
          _436_v36 = _dafny.Map.Empty.slice().updateUnsafe(_415_v28,false);
          let _437_v37;
          _437_v37 = _dafny.MultiSet.fromElements((((_436_v36).contains(_415_v28)) ? ((_436_v36).get(_415_v28)) : (_364_v1)));
          let _438_v38;
          _438_v38 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,_this.f18);
          let _index72 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_426_cf5).length));
          let _rhs71 = (_431_v33).equals(((_431_v33).update((_434_v34).update(p2, _module.__default.abs(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_364_v1),_364_v1)).length))), _module.__default.abs(p2))).Union(_431_v33));
          let _rhs72 = (_437_v37).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_435_v35, _435_v35)));
          let _rhs73 = _dafny.Seq.Concat((_418_v30).fm7(_364_v1, _this.f18, _438_v38, p1, globalState), p1);
          let _rhs74 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_427_cf4, _430_cf1));
          let _lhs55 = _426_cf5;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_426_cf5).length));
          let _lhs57 = globalState;
          let _lhs58 = globalState;
          _lhs55[_lhs56] = _rhs71;
          _364_v1 = _rhs72;
          _lhs57.f3 = _rhs73;
          _lhs58.f10 = _rhs74;
          (globalState).f3 = p1;
          let _439_v39;
          _439_v39 = _module.D0.create_DC0(new BigNumber((_438_v38).length));
          let _440_v40;
          _440_v40 = _dafny.Seq.of(_439_v39);
          _440_v40 = _dafny.Seq.Concat(_module.__default.fm10(_364_v1, _415_v28, globalState), _440_v40);
        } else if (_source3.is_DC2) {
          let _441___mcc_h5 = (_source3).cf6;
          let _442_cf6 = _441___mcc_h5;
          let _rhs75 = false;
          let _rhs76 = (_module.__default.fm0(globalState)).minus(_this.f18);
          let _rhs77 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p2, new BigNumber((p1).length)), _this.f18);
          let _lhs59 = globalState;
          let _lhs60 = globalState;
          _364_v1 = _rhs75;
          _lhs59.f11 = _rhs76;
          _lhs60.f10 = _rhs77;
          let _443_v41;
          _443_v41 = _dafny.Map.Empty.slice().updateUnsafe(_442_cf6,_442_cf6);
          _443_v41 = ((_443_v41).update(_364_v1, _442_cf6)).Merge(_443_v41);
          let _index73 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((p0).length));
          (p0)[_index73] = _442_cf6;
          let _index74 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((p0).length));
          (p0)[_index74] = false;
          let _444_v42;
          _444_v42 = _dafny.Seq.of((p0)[_module.__default.safeIndex(new BigNumber(647), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(647), new BigNumber((p0).length))], _442_cf6);
          _444_v42 = _444_v42;
        } else if (_source3.is_DC3) {
          let _445___mcc_h6 = (_source3).cf7;
          let _446___mcc_h7 = (_source3).cf8;
          let _447___mcc_h8 = (_source3).cf9;
          let _448___mcc_h9 = (_source3).cf10;
          let _449_cf10 = _448___mcc_h9;
          let _450_cf9 = _447___mcc_h8;
          let _451_cf8 = _446___mcc_h7;
          let _452_cf7 = _445___mcc_h6;
          let _index75 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length));
          (p0)[_index75] = _364_v1;
          let _index76 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length));
          (p0)[_index76] = !(_364_v1) || (_364_v1);
          _451_cf8 = (_dafny.MultiSet.FromArray(_module.__default.fm4(globalState))).IsDisjointFrom(_dafny.MultiSet.fromElements(_451_cf8));
          let _453_v43;
          _453_v43 = _module.D4.create_DC9(_418_v30);
          let _454_v44;
          _454_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_451_cf8, _451_cf8, _this.f18, globalState),_418_v30);
          let _455_v45;
          let _nw75 = Array((new BigNumber(17)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _418_v30;
          _nw75[(_dafny.ONE).toNumber()] = _418_v30;
          _nw75[(new BigNumber(2)).toNumber()] = (_453_v43).dtor_cf21;
          _nw75[(new BigNumber(3)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(4)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(5)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(6)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(7)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(8)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(9)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(10)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(11)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(12)).toNumber()] = ((_364_v1) ? (_418_v30) : (_418_v30));
          _nw75[(new BigNumber(13)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(14)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(15)).toNumber()] = _418_v30;
          _nw75[(new BigNumber(16)).toNumber()] = (((_454_v44).contains((p0)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length))])) ? ((_454_v44).get((p0)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length))])) : (_418_v30));
          _455_v45 = _nw75;
          let _index77 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_455_v45).length));
          (_455_v45)[_index77] = _418_v30;
          let _index78 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length));
          let _index79 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_455_v45).length));
          let _rhs78 = (_451_cf8) || (_451_cf8);
          let _rhs79 = _418_v30;
          let _lhs61 = p0;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((p0).length));
          let _lhs63 = _455_v45;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_455_v45).length));
          _lhs61[_lhs62] = _rhs78;
          _lhs63[_lhs64] = _rhs79;
          (globalState).f10 = _452_cf7;
        } else {
          let _456___mcc_h10 = (_source3).cf0;
          let _457_cf0 = _456___mcc_h10;
          let _458_v46;
          _458_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_364_v1);
          let _459_v48;
          _459_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(305),function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(306), new BigNumber(513))) {
              let _460_v47 = _compr_16;
              if (((new BigNumber(306)).isLessThanOrEqualTo(_460_v47)) && ((_460_v47).isLessThan(new BigNumber(513)))) {
                _coll16.push([(_460_v47).minus(_457_cf0),_364_v1]);
              }
            }
            return _coll16;
          }());
          let _461_v49;
          _461_v49 = _module.D0.create_DC3(p2, _364_v1, _this.f18, _this.f18);
          let _462_v50;
          _462_v50 = _dafny.Map.Empty.slice().updateUnsafe((_458_v46).Merge((((_459_v48).contains(_457_cf0)) ? ((_459_v48).get(_457_cf0)) : (_458_v46))),_461_v49);
          _462_v50 = (_462_v50).update(_458_v46, _461_v49);
          let _463_v51;
          _463_v51 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,_this.f18);
          let _464_v52;
          _464_v52 = p1;
          (globalState).f3 = (_418_v30).fm7(_364_v1, _this.f18, ((_364_v1) ? (_463_v51) : (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_364_v1, true, _457_cf0, globalState),new BigNumber(175)))), (_464_v52), globalState);
          (globalState).f12 = _364_v1;
          (_this).f18 = (new BigNumber(252)).minus(p2);
        }
        (globalState).f12 = false;
      } else {
        let _465_v53;
        _465_v53 = _module.D1.create_DC5(_364_v1, _this.f18);
        let _source4 = _465_v53;
        if (_source4.is_DC5) {
          let _466___mcc_h11 = (_source4).cf12;
          let _467___mcc_h12 = (_source4).cf13;
          let _468_cf13 = _467___mcc_h12;
          let _469_cf12 = _466___mcc_h11;
          let _470_v54;
          _470_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(528),_469_cf12);
          let _471_v55;
          _471_v55 = _dafny.Seq.of(_469_cf12, _469_cf12);
          let _472_v56;
          _472_v56 = new _dafny.CodePoint('a'.codePointAt(0));
          let _473_v57;
          _473_v57 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_471_v55, (((_470_v54).contains(_468_cf13)) ? ((_470_v54).get(_468_cf13)) : (_469_cf12))),_dafny.Map.Empty.slice().updateUnsafe(p2,_472_v56));
          _473_v57 = (_473_v57).Merge(_473_v57);
          let _474_v58;
          let _nw76 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _474_v58 = _nw76;
          let _475_v59;
          _475_v59 = _474_v58;
          let _rhs80 = (_475_v59);
          let _rhs81 = p2;
          let _lhs65 = _this;
          _474_v58 = _rhs80;
          _lhs65.f18 = _rhs81;
          let _476_v60;
          _476_v60 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,((_dafny.ZERO).minus(new BigNumber((_module.__default.fm11(_364_v1, globalState)).length))).plus(p2));
          _476_v60 = (_476_v60).update(new BigNumber(-983), p2);
          (globalState).f10 = new BigNumber(-458);
        } else {
          let _477___mcc_h13 = (_source4).cf11;
          let _478_cf11 = _477___mcc_h13;
          let _479_v61;
          let _nw77 = new _module.C0();
          _nw77.__ctor();
          _479_v61 = _nw77;
          _479_v61 = _479_v61;
          _364_v1 = _364_v1;
          _364_v1 = _364_v1;
          let _index80 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((p0).length));
          (p0)[_index80] = (_this.f18).isLessThanOrEqualTo(_this.f18);
          let _index81 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((p0).length));
          (p0)[_index81] = _364_v1;
        }
        let _480_v63;
        _480_v63 = _dafny.Set.fromElements(_364_v1);
        let _481_v64;
        let _nw78 = new _module.C0();
        _nw78.__ctor();
        _481_v64 = _nw78;
        let _482_v65;
        _482_v65 = _dafny.Seq.of(_this.f18, new BigNumber(55));
        let _483_v66;
        _483_v66 = _dafny.Seq.of(new BigNumber((function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(905), new BigNumber(243))) {
            let _484_v62 = _compr_17;
            if (((new BigNumber(905)).isLessThanOrEqualTo(_484_v62)) && ((_484_v62).isLessThan(new BigNumber(243)))) {
              _coll17.push([(_484_v62).multipliedBy(p2),_364_v1]);
            }
          }
          return _coll17;
        }()).length), new BigNumber((_480_v63).length), new BigNumber((_dafny.Seq.of(_481_v64)).length), new BigNumber((_482_v65).length));
        (globalState).f12 = _module.__default.fm1(_364_v1, false, new BigNumber((_483_v66).length), globalState);
        (globalState).f12 = false;
        let _485_v67;
        let _nw79 = Array((new BigNumber(25)).toNumber()).fill([]);
        _485_v67 = _nw79;
        let _486_v68;
        let _init11 = function (_487_i17) {
          return _dafny.MultiSet.fromElements(_this.f18);
        };
        let _nw80 = Array((new BigNumber(2)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw80.length); _i0_11++) {
          _nw80[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _486_v68 = _nw80;
        let _index82 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_485_v67).length));
        (_485_v67)[_index82] = _486_v68;
        let _488_v69;
        _488_v69 = _dafny.MultiSet.fromElements(_364_v1);
        let _489_v70;
        _489_v70 = _dafny.MultiSet.fromElements(_this.f18, p2, (_dafny.ZERO).minus(p2), new BigNumber((_488_v69).cardinality()));
        let _490_v71;
        _490_v71 = _dafny.Seq.of(_489_v70);
        let _index83 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_485_v67).length));
        let _nw81 = Array((new BigNumber(15)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = (_490_v71)[_module.__default.safeIndex(p2, new BigNumber((_490_v71).length))];
        _nw81[(_dafny.ONE).toNumber()] = ((_481_v64).fm6(new BigNumber((_483_v66).length), _364_v1, globalState)).Difference(_489_v70);
        _nw81[(new BigNumber(2)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_483_v66, _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_483_v66).length)), new BigNumber(706)), _483_v66));
        _nw81[(new BigNumber(4)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(5)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(6)).toNumber()] = (_489_v70).Intersect(_489_v70);
        _nw81[(new BigNumber(7)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(8)).toNumber()] = (_489_v70).update(p2, _module.__default.abs(p2));
        _nw81[(new BigNumber(9)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(10)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(11)).toNumber()] = (_489_v70).Difference(_489_v70);
        _nw81[(new BigNumber(12)).toNumber()] = _489_v70;
        _nw81[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(p2, _this.f18);
        _nw81[(new BigNumber(14)).toNumber()] = (_489_v70).Union((_dafny.MultiSet.fromElements(p2, _this.f18)).update(p2, _module.__default.abs(_this.f18)));
        (_485_v67)[_index83] = _nw81;
        let _491_v72;
        _491_v72 = _dafny.Seq.of(p0, p0);
        _491_v72 = _491_v72;
      }
      if ((_364_v1) || (!(_364_v1) || (_364_v1))) {
        let _492_v73;
        _492_v73 = _dafny.Set.fromElements(true);
        let _493_v74;
        _493_v74 = new _dafny.CodePoint('i'.codePointAt(0));
        let _494_v75;
        _494_v75 = _dafny.Seq.of(_492_v73, _module.__default.fm12(_493_v74, _this.f18, globalState));
        let _495_v76;
        _495_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_494_v75)[_module.__default.safeIndex(p2, new BigNumber((_494_v75).length))]).length),_363_v0);
        let _496_v77;
        _496_v77 = _module.D6.create_DC13(_495_v76);
        let _source5 = _module.D0.create_DC3(p2, _module.__default.fm1(_364_v1, _364_v1, new BigNumber(((_496_v77).dtor_cf27).length), globalState), p2, p2);
        if (_source5.is_DC1) {
          let _497___mcc_h14 = (_source5).cf1;
          let _498___mcc_h15 = (_source5).cf2;
          let _499___mcc_h16 = (_source5).cf3;
          let _500___mcc_h17 = (_source5).cf4;
          let _501___mcc_h18 = (_source5).cf5;
          let _502_cf5 = _501___mcc_h18;
          let _503_cf4 = _500___mcc_h17;
          let _504_cf3 = _499___mcc_h16;
          let _505_cf2 = _498___mcc_h15;
          let _506_cf1 = _497___mcc_h14;
          let _index84 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_363_v0).length));
          (_363_v0)[_index84] = p2;
          let _507_v78;
          _507_v78 = _dafny.MultiSet.fromElements(_364_v1, _364_v1);
          let _508_v79;
          _508_v79 = _dafny.Seq.of(_364_v1, _364_v1, _364_v1, _364_v1, _364_v1);
          let _index85 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_363_v0).length));
          (_363_v0)[_index85] = ((_505_cf2).minus((((_507_v78).contains(_364_v1)) ? ((_507_v78).get(_364_v1)) : (p2)))).multipliedBy(new BigNumber((_module.__default.fm13(_364_v1, _508_v79, globalState)).length));
          let _509_v80;
          let _nw82 = new _module.C0();
          _nw82.__ctor();
          _509_v80 = _nw82;
          let _510_v81;
          let _nw83 = new _module.C0();
          _nw83.__ctor();
          _510_v81 = _nw83;
          let _511_v82;
          _511_v82 = _dafny.Map.Empty.slice().updateUnsafe(_509_v80,_505_cf2);
          let _512_v83;
          _512_v83 = _module.D3.create_DC7(_511_v82);
          let _pat_let_tv5 = _511_v82;
          _512_v83 = function (_pat_let7_0) {
            return function (_513_dt__update__tmp_h1) {
              return function (_pat_let8_0) {
                return function (_514_dt__update_hcf15_h0) {
                  return _module.D3.create_DC7(_514_dt__update_hcf15_h0);
                }(_pat_let8_0);
              }(_pat_let_tv5);
            }(_pat_let7_0);
          }(((_364_v1) ? (_512_v83) : (_512_v83)));
        } else if (_source5.is_DC2) {
          let _515___mcc_h19 = (_source5).cf6;
          let _516_cf6 = _515___mcc_h19;
          let _517_v84;
          let _nw84 = new _module.C0();
          _nw84.__ctor();
          _517_v84 = _nw84;
          let _518_v85;
          _518_v85 = _dafny.Map.Empty.slice().updateUnsafe(!(_516_cf6),_517_v84);
          _518_v85 = (_518_v85).update(_364_v1, _517_v84);
          let _index86 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length));
          (p0)[_index86] = _516_cf6;
          let _519_v86;
          _519_v86 = _dafny.Map.Empty.slice().updateUnsafe(_516_cf6,_364_v1);
          let _520_v87;
          _520_v87 = _dafny.Seq.of(p2, new BigNumber((_519_v86).length));
          let _521_v88;
          _521_v88 = _dafny.MultiSet.fromElements(p2);
          let _index87 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length));
          (p0)[_index87] = ((_dafny.MultiSet.FromArray(_520_v87)).Intersect(_521_v88)).IsProperSubsetOf(_521_v88);
          let _522_v89;
          _522_v89 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), ((_523_p2) => function (_524_i18) {
            return _523_p2;
          })(p2)),_this.f18);
          _522_v89 = (_522_v89).update(_520_v87, (p2).minus(p2));
          let _525_v90;
          let _nw85 = Array((new BigNumber(18)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(_dafny.ONE).toNumber()] = !(_364_v1);
          _nw85[(new BigNumber(2)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(new BigNumber(3)).toNumber()] = false;
          _nw85[(new BigNumber(4)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(new BigNumber(5)).toNumber()] = false;
          _nw85[(new BigNumber(6)).toNumber()] = false;
          _nw85[(new BigNumber(7)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(8)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(9)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(new BigNumber(10)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(new BigNumber(11)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(12)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((p0).length))];
          _nw85[(new BigNumber(13)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(14)).toNumber()] = _364_v1;
          _nw85[(new BigNumber(15)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(16)).toNumber()] = _516_cf6;
          _nw85[(new BigNumber(17)).toNumber()] = true;
          _525_v90 = _nw85;
          let _526_v91;
          _526_v91 = _dafny.MultiSet.fromElements(_525_v90);
          (_this).f18 = new BigNumber((_dafny.Seq.Concat(_520_v87, _dafny.Seq.of(p2, new BigNumber((_526_v91).cardinality())))).length);
        } else if (_source5.is_DC3) {
          let _527___mcc_h20 = (_source5).cf7;
          let _528___mcc_h21 = (_source5).cf8;
          let _529___mcc_h22 = (_source5).cf9;
          let _530___mcc_h23 = (_source5).cf10;
          let _531_cf10 = _530___mcc_h23;
          let _532_cf9 = _529___mcc_h22;
          let _533_cf8 = _528___mcc_h21;
          let _534_cf7 = _527___mcc_h20;
          let _535_v92;
          _535_v92 = _dafny.Seq.of(_364_v1, _364_v1);
          let _536_v93;
          _536_v93 = _dafny.MultiSet.fromElements(_module.__default.fm1(_364_v1, _533_cf8, _this.f18, globalState), _364_v1, _364_v1, _364_v1, _364_v1);
          let _537_v94;
          _537_v94 = _dafny.Seq.of((p2).plus(new BigNumber((_535_v92).length)), (((_536_v93).contains(_533_cf8)) ? ((_536_v93).get(_533_cf8)) : (p2)), _531_cf10);
          _537_v94 = ((_364_v1) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), ((_538_cf7) => function (_539_i19) {
            return new BigNumber((_dafny.Set.fromElements(_538_cf7)).length);
          })(_534_cf7))) : (_537_v94));
          let _540_v95;
          _540_v95 = _dafny.MultiSet.fromElements(_493_v74);
          let _541_v96;
          _541_v96 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_module.__default.fm0(globalState), p2),_540_v95);
          _541_v96 = (_541_v96).update(_532_cf9, _540_v95);
          let _nw86 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _363_v0 = _nw86;
          (globalState).f10 = _this.f18;
        } else {
          let _542___mcc_h24 = (_source5).cf0;
          let _543_cf0 = _542___mcc_h24;
          let _544_v97;
          _544_v97 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,_364_v1);
          let _rhs82 = _module.__default.fm0(globalState);
          let _rhs83 = (new BigNumber(((_544_v97).Merge(_544_v97)).length)).isLessThanOrEqualTo(_543_cf0);
          let _rhs84 = _363_v0;
          let _lhs66 = globalState;
          let _lhs67 = globalState;
          _lhs66.f11 = _rhs82;
          _lhs67.f12 = _rhs83;
          _363_v0 = _rhs84;
          let _545_v98;
          _545_v98 = _dafny.MultiSet.fromElements(_364_v1);
          _545_v98 = ((_545_v98).Difference(_545_v98)).Union(_module.__default.fm14(_this.f18, _module.D1.create_DC5(_364_v1, _543_cf0), globalState));
          (globalState).f12 = _364_v1;
          let _546_v99;
          let _nw87 = new _module.C0();
          _nw87.__ctor();
          _546_v99 = _nw87;
        }
        _363_v0 = _363_v0;
        let _547_v100;
        _547_v100 = _dafny.Set.fromElements(p2);
        _547_v100 = _547_v100;
        _364_v1 = _dafny.areEqual(p1, p1);
        let _548_v101;
        _548_v101 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,_this.f18);
        let _549_v102;
        _549_v102 = _module.D0.create_DC3(new BigNumber(429), _364_v1, new BigNumber((_548_v101).length), new BigNumber(366));
        let _550_v103;
        _550_v103 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_549_v102).dtor_cf8);
        _550_v103 = _550_v103;
      } else {
        (globalState).f10 = _this.f18;
        let _551_v104;
        let _init12 = function (_552_i20) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        };
        let _nw88 = Array((new BigNumber(11)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw88.length); _i0_12++) {
          _nw88[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _551_v104 = _nw88;
        _551_v104 = _551_v104;
        let _553_v105;
        _553_v105 = new _dafny.CodePoint('l'.codePointAt(0));
        let _554_v106;
        _554_v106 = _dafny.Set.fromElements(_553_v105, _module.__default.fm9(_364_v1, !(false), globalState), _553_v105, _553_v105);
        let _555_v107;
        _555_v107 = _module.D0.create_DC2((new BigNumber((_554_v106).length)).isLessThanOrEqualTo(p2));
        let _source6 = _555_v107;
        if (_source6.is_DC1) {
          let _556___mcc_h25 = (_source6).cf1;
          let _557___mcc_h26 = (_source6).cf2;
          let _558___mcc_h27 = (_source6).cf3;
          let _559___mcc_h28 = (_source6).cf4;
          let _560___mcc_h29 = (_source6).cf5;
          let _561_cf5 = _560___mcc_h29;
          let _562_cf4 = _559___mcc_h28;
          let _563_cf3 = _558___mcc_h27;
          let _564_cf2 = _557___mcc_h26;
          let _565_cf1 = _556___mcc_h25;
          let _index88 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_561_cf5).length));
          (_561_cf5)[_index88] = (_module.__default.fm1(_364_v1, _364_v1, _module.__default.fm0(globalState), globalState)) && (_364_v1);
          let _566_v108;
          _566_v108 = _dafny.Seq.of(_565_cf1);
          let _567_v109;
          _567_v109 = _dafny.Map.Empty.slice().updateUnsafe(_363_v0,_364_v1);
          let _568_v110;
          _568_v110 = _dafny.MultiSet.fromElements(new BigNumber(786));
          let _index89 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_561_cf5).length));
          (_561_cf5)[_index89] = ((_dafny.MultiSet.FromArray(_566_v108)).update(_563_cf3, _module.__default.abs(new BigNumber((_567_v109).length)))).IsDisjointFrom(_568_v110);
          let _569_v111;
          _569_v111 = _dafny.Seq.of((_561_cf5)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_561_cf5).length))]);
          _569_v111 = _569_v111;
          let _570_v112;
          let _nw89 = new _module.C0();
          _nw89.__ctor();
          _570_v112 = _nw89;
          _364_v1 = (_569_v111)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_569_v111).length))];
        } else if (_source6.is_DC2) {
          let _571___mcc_h30 = (_source6).cf6;
          let _572_cf6 = _571___mcc_h30;
          let _573_v113;
          _573_v113 = _module.D6.create_DC14(false, _553_v105, _551_v104);
          _573_v113 = _573_v113;
          let _574_v114;
          let _nw90 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _574_v114 = _nw90;
          let _575_v115;
          _575_v115 = _module.D4.create_DC10(_553_v105, _574_v114, _this.f18);
          let _576_v116;
          _576_v116 = _dafny.Set.fromElements(_575_v115);
          let _577_v117;
          _577_v117 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_575_v115)).IsSubsetOf(_576_v116),_364_v1);
          let _578_v118;
          let _init13 = ((_579_v107) => function (_580_i21) {
            return _579_v107;
          })(_555_v107);
          let _nw91 = Array((new BigNumber(7)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw91.length); _i0_13++) {
            _nw91[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _578_v118 = _nw91;
          let _index90 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_578_v118).length));
          (_578_v118)[_index90] = _555_v107;
          let _index91 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_578_v118).length));
          let _rhs85 = _577_v117;
          let _rhs86 = _dafny.Seq.of(_553_v105);
          let _rhs87 = _555_v107;
          let _rhs88 = _572_cf6;
          let _lhs68 = globalState;
          let _lhs69 = _578_v118;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_578_v118).length));
          let _lhs71 = globalState;
          _577_v117 = _rhs85;
          _lhs68.f3 = _rhs86;
          _lhs69[_lhs70] = _rhs87;
          _lhs71.f12 = _rhs88;
          let _581_v119;
          let _nw92 = new _module.C0();
          _nw92.__ctor();
          _581_v119 = _nw92;
          let _582_v120;
          let _nw93 = new _module.C0();
          _nw93.__ctor();
          _582_v120 = _nw93;
        } else if (_source6.is_DC3) {
          let _583___mcc_h31 = (_source6).cf7;
          let _584___mcc_h32 = (_source6).cf8;
          let _585___mcc_h33 = (_source6).cf9;
          let _586___mcc_h34 = (_source6).cf10;
          let _587_cf10 = _586___mcc_h34;
          let _588_cf9 = _585___mcc_h33;
          let _589_cf8 = _584___mcc_h32;
          let _590_cf7 = _583___mcc_h31;
          let _591_v121;
          let _nw94 = new _module.C0();
          _nw94.__ctor();
          _591_v121 = _nw94;
          let _592_v122;
          _592_v122 = _dafny.Map.Empty.slice().updateUnsafe(_553_v105,_590_cf7);
          (globalState).f11 = new BigNumber(((_592_v122).Merge(_592_v122)).length);
          let _593_v123;
          let _nw95 = new _module.C0();
          _nw95.__ctor();
          _593_v123 = _nw95;
          let _594_v124;
          let _nw96 = Array((new BigNumber(28)).toNumber());
          _594_v124 = _nw96;
          let _index92 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_594_v124).length));
          (_594_v124)[_index92] = _591_v121;
          let _index93 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_594_v124).length));
          (_594_v124)[_index93] = _593_v123;
        } else {
          let _595___mcc_h35 = (_source6).cf0;
          let _596_cf0 = _595___mcc_h35;
          let _597_v125;
          _597_v125 = _dafny.Set.fromElements(_this.f18);
          (globalState).f12 = (_dafny.Set.fromElements(_596_cf0, _596_cf0)).IsSubsetOf(_597_v125);
          let _598_v126;
          _598_v126 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,p2);
          let _599_v127;
          _599_v127 = _dafny.Seq.of(_364_v1);
          let _600_v131;
          let _nw97 = new _module.C0();
          _nw97.__ctor();
          _600_v131 = _nw97;
          let _601_v132;
          _601_v132 = _dafny.Set.fromElements(_600_v131);
          let _602_v133;
          let _nw98 = Array((new BigNumber(21)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = (_598_v126).Merge(_598_v126);
          _nw98[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_this.f18);
          _nw98[(new BigNumber(2)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(3)).toNumber()] = (((_599_v127)[_module.__default.safeIndex(_this.f18, new BigNumber((_599_v127).length))]) ? (_598_v126) : (_598_v126));
          _nw98[(new BigNumber(4)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(5)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(6)).toNumber()] = function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(241), new BigNumber(378))) {
              let _603_v128 = _compr_18;
              if (((new BigNumber(241)).isLessThanOrEqualTo(_603_v128)) && ((_603_v128).isLessThan(new BigNumber(378)))) {
                _coll18.push([_module.__default.safeModuloInt(_603_v128, (_dafny.ZERO).minus(new BigNumber((_599_v127).length))),_this.f18]);
              }
            }
            return _coll18;
          }();
          _nw98[(new BigNumber(7)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(8)).toNumber()] = (function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_598_v126).Keys.Elements) {
              let _604_v129 = _compr_19;
              if ((_598_v126).contains(_604_v129)) {
                _coll19.push([(_604_v129).plus((_dafny.ZERO).minus(p2)),_this.f18]);
              }
            }
            return _coll19;
          }()).Merge(function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(735), new BigNumber(507))) {
              let _605_v130 = _compr_20;
              if (((new BigNumber(735)).isLessThanOrEqualTo(_605_v130)) && ((_605_v130).isLessThan(new BigNumber(507)))) {
                _coll20.push([_module.__default.safeDivisionInt(_605_v130, _596_cf0),new BigNumber((_dafny.MultiSet.fromElements(_596_cf0)).cardinality())]);
              }
            }
            return _coll20;
          }());
          _nw98[(new BigNumber(9)).toNumber()] = _module.__default.fm15(new BigNumber((_601_v132).length), _364_v1, _596_cf0, globalState);
          _nw98[(new BigNumber(10)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(11)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_596_cf0,new BigNumber(650));
          _nw98[(new BigNumber(13)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(14)).toNumber()] = (_598_v126).Merge(_598_v126);
          _nw98[(new BigNumber(15)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_596_cf0);
          _nw98[(new BigNumber(17)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(18)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(19)).toNumber()] = _598_v126;
          _nw98[(new BigNumber(20)).toNumber()] = _598_v126;
          _602_v133 = _nw98;
          let _nw99 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _602_v133 = _nw99;
          let _606_v134;
          let _nw100 = Array((new BigNumber(14)).toNumber()).fill(_module.D4.Default());
          _606_v134 = _nw100;
          let _607_v135;
          let _nw101 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
          _607_v135 = _nw101;
          let _608_v136;
          _608_v136 = _module.D4.create_DC10(_553_v105, _607_v135, (_dafny.ZERO).minus(_this.f18));
          let _index94 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_606_v134).length));
          (_606_v134)[_index94] = _608_v136;
          let _index95 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_606_v134).length));
          (_606_v134)[_index95] = _608_v136;
          let _609_v137;
          _609_v137 = _dafny.MultiSet.fromElements(p1);
          _598_v126 = (_598_v126).update(new BigNumber((_609_v137).cardinality()), p2);
        }
        let _610_v138;
        _610_v138 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,_364_v1);
        let _611_v139;
        let _init14 = ((_612_v1) => function (_613_i22) {
          return _dafny.Set.fromElements(_612_v1);
        })(_364_v1);
        let _nw102 = Array((new BigNumber(9)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw102.length); _i0_14++) {
          _nw102[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _611_v139 = _nw102;
        let _614_v140;
        _614_v140 = _611_v139;
        let _source7 = (((((_610_v138).contains(_364_v1)) ? ((_610_v138).get(_364_v1)) : (_364_v1))) ? (_611_v139) : (_614_v140));
        let _615___mcc_h36 = _source7;
        let _616_cf26 = _615___mcc_h36;
        let _index96 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((p0).length));
        (p0)[_index96] = !(_this.f18).isEqualTo(p2);
        let _index97 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((p0).length));
        (p0)[_index97] = _364_v1;
        let _617_v141;
        let _nw103 = new _module.C0();
        _nw103.__ctor();
        _617_v141 = _nw103;
        (globalState).f11 = _this.f18;
        let _618_v142;
        _618_v142 = _dafny.Seq.of(_364_v1, _364_v1, (p0)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((p0).length))]);
        let _619_v143;
        _619_v143 = _dafny.Seq.of(_551_v104);
        let _620_v144;
        _620_v144 = _module.D6.create_DC14(false, _553_v105, (_619_v143)[_module.__default.safeIndex(p2, new BigNumber((_619_v143).length))]);
        let _621_v145;
        _621_v145 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_618_v142, _618_v142),(((p0)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((p0).length))]) ? (true) : ((_620_v144).dtor_cf28)));
        _621_v145 = (_621_v145).update(_dafny.Seq.of(_364_v1), !(_364_v1));
        if (!(_364_v1) || (!(((_364_v1) ? (_364_v1) : (_364_v1))))) {
          let _622_v146;
          let _nw104 = new _module.C0();
          _nw104.__ctor();
          _622_v146 = _nw104;
          let _623_v147;
          let _out12;
          _out12 = _module.__default.m0(p2, globalState);
          _623_v147 = _out12;
          let _index98 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_363_v0).length));
          (_363_v0)[_index98] = ((_module.__default.fm1(!(_364_v1), true, _this.f18, globalState)) ? (_this.f18) : (_this.f18));
          let _index99 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_363_v0).length));
          (_363_v0)[_index99] = (((_dafny.ZERO).minus(_this.f18)).minus(_this.f18)).multipliedBy((_dafny.ZERO).minus((p2).multipliedBy(_this.f18)));
          let _624_v148;
          _624_v148 = _dafny.Map.Empty.slice().updateUnsafe(_364_v1,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cudbn")).length)));
          (globalState).f3 = (_622_v146).fm7(_364_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), ((_625_v105) => function (_626_i23) {
            return _625_v105;
          })(_553_v105))).length), _624_v148, p1, globalState);
          let _rhs89 = _553_v105;
          let _rhs90 = _this.f18;
          let _lhs72 = globalState;
          _553_v105 = _rhs89;
          _lhs72.f10 = _rhs90;
        } else {
          let _627_v149;
          let _nw105 = new _module.C0();
          _nw105.__ctor();
          _627_v149 = _nw105;
          let _628_v151;
          let _init15 = ((_629_p2) => function (_630_i24) {
            return _dafny.Seq.of(_this.f18, new BigNumber((function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of _dafny.IntegerRange(new BigNumber(640), new BigNumber(-60))) {
                let _631_v150 = _compr_21;
                if (((new BigNumber(640)).isLessThanOrEqualTo(_631_v150)) && ((_631_v150).isLessThan(new BigNumber(-60)))) {
                  _coll21.push([(_631_v150).multipliedBy(_629_p2),_this.f18]);
                }
              }
              return _coll21;
            }()).length));
          })(p2);
          let _nw106 = Array((new BigNumber(4)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw106.length); _i0_15++) {
            _nw106[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _628_v151 = _nw106;
          let _632_v152;
          _632_v152 = _module.D0.create_DC3(_this.f18, false, _this.f18, new BigNumber((_554_v106).length));
          let _633_v153;
          _633_v153 = _module.D4.create_DC10(_module.__default.fm9(_364_v1, _364_v1, globalState), _628_v151, new BigNumber((_dafny.Set.fromElements(_632_v152)).length));
          _633_v153 = _633_v153;
          let _634_v154;
          _634_v154 = _dafny.Seq.of(_363_v0, _363_v0);
          _364_v1 = !(_module.__default.fm1(_364_v1, _dafny.Seq.contains(_634_v154, _363_v0), (p2).plus(p2), globalState));
          (globalState).f10 = ((_364_v1) ? (p2) : ((_this.f18).minus(p2)));
          (globalState).f10 = _this.f18;
        }
      }
      let _635_v155;
      _635_v155 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),!(_364_v1));
      let _636_v156;
      _636_v156 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((_364_v1) ? (_635_v155) : (_635_v155))).length),p2);
      _636_v156 = (_636_v156).update((_dafny.ZERO).minus(_this.f18), (_dafny.ZERO).minus(new BigNumber((p1).length)));
      return;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
